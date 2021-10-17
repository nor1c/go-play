package main

import (
	"fmt"
)

func main() {
	var i interface{}
	describe(i) // nil

	i = 42
	describe(i) // int

	i = "a str"
	describe(i) // string
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T\n)", i, i)
}
