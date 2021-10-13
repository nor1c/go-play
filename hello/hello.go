package main

import (
	"fmt"

	"github.com/nor1c/go-play/greetings"
)

func main() {
	message := greetings.Hello("Fauzi")
	fmt.Println(message)
}
