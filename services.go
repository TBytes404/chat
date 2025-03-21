package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func serve(handler func(net.Conn)) {
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
		go handler(con)
	}
}

func connect(handler func(net.Conn)) {
	con, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	fmt.Println("Client Connected to Post 8000...")
	handler(con)
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
