package main

import (
	"fmt"
	"log"
	"net"
	"slices"
)

type User struct {
	Name string
	Conn net.Conn
}

func NewUser(name string, conn net.Conn) *User {
	return &User{name, conn}
}

func (u *User) Receive(text string) {
	if _, err := fmt.Fprintln(u.Conn, text); err != nil {
		log.Fatal(err)
	}
}

type Message struct {
	Text   string
	Author *User
}

func NewMessage(author *User, text string) *Message {
	return &Message{text, author}
}

type Room struct {
	Name     string
	Members  []*User
	Messages []*Message
}

func NewRoom(name string) *Room {
	return &Room{name, []*User{}, []*Message{}}
}

func (r *Room) Subscribe(user *User) {
	r.Members = append(r.Members, user)
}

func (r *Room) Unsubscribe(user *User) {
	i := slices.Index(r.Members, user)
	if i != -1 {
		r.Members = slices.Delete(r.Members, i, i+1)
	}
}

func (r *Room) Broadcast(msg *Message) {
	text := ""
	if msg != nil {
		author := "[!] Server"
		if msg.Author != nil {
			author = msg.Author.Name
		}
		text = author + " : " + msg.Text
	}
	for _, u := range r.Members {
		u.Receive(text)
	}
}

func (r *Room) Publish(msg *Message) {
	if slices.Contains(r.Members, msg.Author) {
		r.Messages = append(r.Messages, msg)
	}
	r.Broadcast(msg)
}
