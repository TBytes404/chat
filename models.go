package main

import (
	"fmt"
	"slices"
)

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{name}
}

func (u *User) Receive(msg *Message) {
	fmt.Printf("To %s,\n%s\n\tby %s\n", u.Name, msg.Text, msg.Author.Name)
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
