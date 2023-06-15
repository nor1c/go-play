package main

import (
	"fmt"
	"math"
	"reflect"
)

const s string = "this is a string"

func main() {
	fmt.Println(s, reflect.TypeOf(s)) // "this is a string" string

	const n = 5000
	fmt.Println(reflect.TypeOf(n)) // int

	const d = 3e20 / n
	fmt.Println(d, reflect.TypeOf(d)) // 6e+16 float64

	fmt.Println(int64(d), reflect.TypeOf(int64(d))) // 60000000000000000, int64

	fmt.Println(math.Sin(n), reflect.TypeOf(math.Sin(n))) // -0.9879.. float64

	const (
		x = 2
		y = 3
	)
	x = 4             // error: cannot assign to x (untyped int constant 2)
	fmt.Println(x, y) // 2 3
}
