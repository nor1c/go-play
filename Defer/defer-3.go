package main

import (
	"fmt"
	"log"
)

func startApp(error bool) error {
	defer endApp()

	fmt.Println("Application started!")

	if error {
		return fmt.Errorf("App error!")
	}

	return nil
}

func endApp() {
	fmt.Println("Application stopped!")
}

func main() {
	err := startApp(true)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("This will not be called!") // not called because `os.Exit(1)` in `log.Fatal()`
}
