package main

import "fmt"

func main() {
	newChan := make(chan string)

	newChan <- "OK!"

	fmt.Println(<-newChan)
}
