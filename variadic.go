package main

import "fmt"

func consume(xs ...int) {
	fmt.Println(xs)
}

func main() {
	xs := []int{1, 2, 3}
	consume(xs...)               // output: [1 2 3]
	consume(append(xs, 4, 5)...) // output: [1 2 3 4 5]
}
