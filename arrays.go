package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println("arr:", arr)

	fmt.Println("length of arr:", len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var a [5]int
	fmt.Println("emp:", a)

	// set value of array index
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// multi-dimensional data structures
	var twoD [2][3]int
	fmt.Println("multiD:", twoD) // [[0 1 2] [1 2 3]]

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
}
