package main

import "fmt"

func squareToChannel(data []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range data {
			out <- n
		}
	}()

	return out
}

func square(data <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range data {
			out <- n * n
		}
	}()

	return out
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	dataCh := squareToChannel(data)

	finalCh := square(dataCh)

	for n := range finalCh {
		fmt.Println(n)
	}
}
