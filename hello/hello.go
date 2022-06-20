package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Ridwan")
	fmt.Println(message)
}
