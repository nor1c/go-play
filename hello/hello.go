package main

import (
	"fmt"
	"log"

	"github.com/nor1c/go-play/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Fauzi")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message) // will return 'empty name' if name is not provided
}
