package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// sets the prefix and the flag
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request the message
	message, err := greetings.Hello("")

	if err != "" {
		log.Fatal(err)
	}

	message := greetings.Hello("Ridwan")
	fmt.Println(message)
}
