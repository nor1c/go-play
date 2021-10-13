package main

import (
	"fmt"
	"log"

	"github.com/nor1c/go-play/greetings"
)

func main() {
	// set properties of the predefined Logger, including the log entry prefix and flag to disable printing the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting message
	name := "Fauzi"
	message, err := greetings.Hello(name)

	// if an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message) // will return 'empty name' if name is not provided
	fmt.Println(err)
}
