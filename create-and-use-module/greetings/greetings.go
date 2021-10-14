package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// if there is no error, return value that embed the name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
