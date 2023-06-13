package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// undefined name error handling
	if name == "" {
		return "", errors.New("Name is required!")
	}

	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

// multiple names
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomGreeting() string {
	formats := []string{
		"Hi, %s. Welcome!",
		"Great to see you, %s!",
		"Hail, %s! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
