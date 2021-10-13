package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!")

	// quote
	fmt.Println(quote.Go())    // Don't communicate by sharing memory, share memory by communicating.
	fmt.Println(quote.Hello()) // Hello, world.
	fmt.Println(quote.Glass()) // I can eat glass and it doesn't hurt me.
	fmt.Println(quote.Opt())   // If a program is too slow, it must have a loop.
}
