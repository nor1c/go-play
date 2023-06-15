package main

import "fmt"

const (
	Big   = 1 << 100  // 1267650600228229401496703205376
	Small = Big >> 99 // 2
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))   // 21
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big))   // 1.2676506002282295e+29
}
