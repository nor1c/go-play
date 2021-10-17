package main

import "fmt"

func main() {
	m := make(map[string]int) // map[key-type]val-type, m[string] = int

	m["k1"] = 7
	m["k2"] = 8
	fmt.Println("map:", m) // map[k1:7 k2:8]

	// get value of map key
	v1 := m["k1"]
	fmt.Println("value of m['k1']:", v1) // 7

	fmt.Println("len:", len(m)) // 2

	// delete a key-value pair
	delete(m, "k2")
	fmt.Println("m after deleted k2:", m) // map[k1:7]

	_, prs := m["k2"]
	fmt.Println("prs:", prs) // false

	_, prs1 := m["k1"]
	fmt.Println("prs1:", prs1) // true

	// declare and initialize map in a single line of code
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n) // map[bar:2 foo:1]
}
