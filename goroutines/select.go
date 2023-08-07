package main

import "fmt"

func main() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	newV := <-myChannel
	fmt.Println(newV) // output: data
}
