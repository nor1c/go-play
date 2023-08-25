package main

import "log"

func transformToChannel(numbers ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range numbers {
			out <- n
		}
	}()

	return out
}

func square(numberCh <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range numberCh {
			select {
			default:
				out <- n * n
			}
		}
	}()

	return out
}

func main() {
	numbersCh := transformToChannel(1, 2, 3)
	sqCh := square(numbersCh)

	for n := range sqCh {
		log.Println(n)
	}

	log.Println("done!")
}
