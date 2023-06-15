package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)

	fmt.Printf("Type of x: %T, y: %T, z: %T\n", x, y, z) // x: int, y: int, z: uint

	var xs = string(x)
	fmt.Printf("Type of xs: %T", xs) // string
}
