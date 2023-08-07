package main

import (
	"fmt"
	"time"
)

/**
GoRoutines is basically an asyunchronous or concurrent process.
You don't need to wait for previous task to finish to run the next process.

Unlike NodeJS, concurrent/asynchronous is run by default. So you need to use `await` to run tasks synchonously/parallel.
In Go, it's the reverse of NodeJS, you need to use keyword `go` to run process asynchonously/concurrently.
**/

func printMe(message string) {
	fmt.Println(message)
}

func main() {
	printMe("1")
	printMe("2")
	printMe("3")
	// until this line, all code processed sequentially, that's mean all following the order

	go printMe("4")
	go printMe("5")
	go printMe("6")
	time.Sleep(time.Second) // wait for 1 second (if this doesn't exist then "4", "5", "6" will never be printed to the console)
	// from printMe("4") to this line, all code processed concurrently, that's mean any code can go out first

	fmt.Println("all processes done!")
}
