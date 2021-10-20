package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) // sum: 9

	// with indexes
	for i, num := range nums {
		fmt.Printf("index of %s: %s", num, i)
		fmt.Println()
	}

	// map key/value pairs
	kvs := map[string]string{"a": "Apple", "b": "Banana", "c": "Coconut"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// a -> Apple
	// b -> Banana
	// c -> Coconut

	// iterate just the key of a map
	for k, v := range "go" {
		fmt.Println(k, v)
	}
	// return an ASCII format:
	// 0 103
	// 1 111
}
