package main

import (
	"fmt"
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

func (u *User) Receive(msg *Message) {
	fmt.Fprintln(u.Conn, msg.Author.Name, ":", msg.Text)
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

func (r *Room) Broadcast(msg *Message) {
	for _, u := range r.Members {
		u.Receive(msg)
	}
}

func (r *Room) Publish(msg *Message) {
	if slices.Contains(r.Members, msg.Author) {
		r.Messages = append(r.Messages, msg)
	}
	r.Broadcast(msg)
}
