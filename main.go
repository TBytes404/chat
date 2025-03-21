package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Your ~Favourite~ Chat Buddy <3")
		return
	}
	switch os.Args[1] {
	case "serve":
		fmt.Println("Server Listening at Post 8000...")
		srv, err := net.Listen("tcp", ":8000")
		if err != nil {
			log.Fatal(err)
		}
		defer srv.Close()
		for {
			con, err := srv.Accept()
			if err != nil {
				log.Fatal(err)
			}
			addr := con.RemoteAddr().String()
			fmt.Println(addr, "Connected...")
			go helper(con, func(msg string) {
				fmt.Println(addr, ":", msg)
				fmt.Fprintln(con, msg)
			})
		}
	case "connect":
		con, err := net.Dial("tcp", ":8000")
		if err != nil {
			log.Fatal(err)
		}
		defer con.Close()
		fmt.Println("Client Connected to Post 8000...")
		go helper(con, func(msg string) { fmt.Println("@", msg) })
		helper(os.Stdin, func(msg string) { fmt.Fprintln(con, msg) })
	}
}

func helper(r io.Reader, fn func(string)) {
	snr := bufio.NewScanner(r)
	for snr.Scan() {
		msg := snr.Text()
		if len(msg) == 0 {
			break
		}
		fn(msg)
	}
	if err := snr.Err(); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
}
