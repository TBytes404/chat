package main

import (
	"bufio"
	"fmt"
	"os"
)

func client() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Username is Invalid!")
		return
	}
	user := NewUser(name)
	fmt.Println(user.Name)
}
