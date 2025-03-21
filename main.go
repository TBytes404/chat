package main

import (
	"fmt"
	"net"
	"os"
)

var room *Room

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Your ~Favourite~ Chat Buddy <3")
		return
	}
	switch os.Args[1] {
	case "serve":
		room = NewRoom("Default")
		serve(func(con net.Conn) {
			user := NewUser(con.RemoteAddr().String(), con)
			room.Subscribe(user)
			fmt.Println(user.Name, "Connected...")
			helper(con, func(msg string) { room.Publish(NewMessage(user, msg)) })
		})
	case "connect":
		connect(func(con net.Conn) {
			go helper(con, func(msg string) { fmt.Println(msg) })
			helper(os.Stdin, func(msg string) { fmt.Fprintln(con, msg) })
		})
	}
}
