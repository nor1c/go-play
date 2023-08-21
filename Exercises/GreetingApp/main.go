package main

import (
	"fmt"
	"log"
	"nor1c/greeting"
)

func main() {
	// greeting Fauzi
	greetFauziMsg, greetFauziErr := greeting.Hello("Fauzi")
	if greetFauziErr != nil {
		log.Fatal(greetFauziErr)
	}
	fmt.Println(greetFauziMsg)

	// greeting undefined name
	log.SetPrefix("[GREETING_LOG] ")
	// log.SetFlags(0)
	greetEmptyMsg, greetEmptyErr := greeting.Hello("")
	if greetEmptyErr != nil {
		log.Println(greetEmptyErr) // Name is required!
	}
	fmt.Println(greetEmptyMsg)

	// multiple greetings
	names := []string{"Fauzi", "John", "Jack"}
	messages, err := greeting.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
