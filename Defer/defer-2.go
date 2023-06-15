package main

import "fmt"

func main() {
	runApp(true)
}

func runApp(isError bool) {
	defer endApp() // it's recommended to put the defer call in the top of the code of function

	if isError {
		panic("[ERR] Application error!")
	}

	fmt.Println("Application running..")
}

func endApp() {
	fmt.Println("Application terminated.")

	// should be called here, because it's no longer running code below `panic()`, but it will still call `endApp()` because it's deferred
	message := recover()
	fmt.Println("Error occured:", message)
}
