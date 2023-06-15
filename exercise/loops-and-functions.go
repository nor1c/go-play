package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1) // decent starting guess for z is 1, no matter what the input
	var prevZ float64

	for {
		z, prevZ = z-(z*z-x)/(2*z), z

		if math.Abs(prevZ-z) < 1e-8 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))

	fmt.Println(1e-8)
}
