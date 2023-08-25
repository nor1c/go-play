package main

import (
	"fmt"

)

func readOnlyChanFunc() <-chan int {
	res := make(chan int)

	go func() {
		res <- 2
	}()

	return res
}

func main() {
	roCh := readOnlyChanFunc()

	fmt.Println(roCh) // 0xc0000...
}
