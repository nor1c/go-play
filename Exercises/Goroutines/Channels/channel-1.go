package main

import "fmt"

func findMax(numbers []int, max chan int) {
	var num int = 0

	for _, v := range numbers {
		if num < v {
			num = v
		}
	}

	max <- num
}

func assignString(message string, strChan chan string) {
	strChan <- message
}

func main() {
	maxChannel := make(chan int)

	numbers := []int{1, 5, 4, 7, 5, 9, 3}
	go findMax(numbers, maxChannel)

	strChan := make(chan string)
	go assignString("Lorem ipsum..", strChan)

	for i := 0; i < 2; i++ {
		select {
		case max := <-maxChannel:
			fmt.Printf("max is %v \n", max)
		case str := <-strChan:
			fmt.Printf("str is %v \n", str)
		}
	}

	fmt.Println("done!")
}
