package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Your ~Favourite~ Chat Buddy <3")
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	switch os.Args[1] {
	case "serve":
		room := NewRoom("Default")

		go func() {
			<-sigs
			fmt.Println("Quiting...")
			room.Broadcast(NewMessage(nil, "Shutdown"))
			room.Broadcast(nil)
			os.Exit(0)
		}()

		serve(func(con net.Conn) {
			user := NewUser(con.RemoteAddr().String(), con)
			room.Subscribe(user)
			fmt.Println(user.Name, "Connected.")

			helper(con, func(msg string) {
				if len(msg) != 0 {
					room.Publish(NewMessage(user, msg))
					return
				}
				room.Unsubscribe(user)
				fmt.Println(user.Name, "Disconnected.")
			})
		})

	case "connect":
		connect(func(con net.Conn) {
			go func() {
				<-sigs
				fmt.Println("Disconnecting...")
				fmt.Fprintln(con, "")
				os.Exit(0)
			}()

			go helper(con, func(msg string) {
				if len(msg) == 0 {
					fmt.Println("Disconnecting...")
					os.Exit(0)
				}
				fmt.Println(msg)
			})

			helper(os.Stdin, func(msg string) {
				if len(msg) != 0 {
					if _, err := fmt.Fprintln(con, msg); err != nil {
						log.Fatal(err)
					}
				}
			})
			fmt.Println("Disconnecting...")
			fmt.Fprintln(con, "")
		})
	}
}
