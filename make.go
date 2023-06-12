package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)

	myMap := make(map[string]string)
	fmt.Println(myMap)

	myChannel := make(chan string)
	fmt.Println(myChannel)
}
