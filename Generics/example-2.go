package main

import "fmt"

func Sum[V Value](v V) V {
	return v
}

func main() {
	fmt.Println(Sum(2))
}
