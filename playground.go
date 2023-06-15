package main

import (
	"fmt"
	"math"
	"math/rand"
)

func calculateMe(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("hello!")
	fmt.Println(math.Pi)
	fmt.Println(rand.Intn(10))

	calculateMe(2, 3)
}
