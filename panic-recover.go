package main

import "fmt"

func iCausePanic() {
	panic("Panicked!")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("App recovered from panic!")
	} else {
		fmt.Println("No panic occured!")
	}
}

func main() {
	defer recoverFromPanic()

	fmt.Println("Hello, world")
	iCausePanic()
	fmt.Println("Hello, guys!")
}
