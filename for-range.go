package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	for _, v := range arr {
		fmt.Println(v)
	} // 1 2 3

	animals := []struct {
		Kind string
		Legs int
	}{
		{"Dog", 4},
		{"Fish", 0},
		{"Chicken", 2},
	}

	for _, v := range animals {
		fmt.Println(v)
	} // {"Dog" 4} {"Fish" 0} {"Chicken" 2}
}
