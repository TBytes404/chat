package main

func example() {
	room := NewRoom("the holi hall")

	user := NewUser("cathe")
	room.Subscribe(user)

	room.Publish(NewMessage(user, "i am an innocent message."))
}
