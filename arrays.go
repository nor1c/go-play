package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println("arr:", arr)

	fmt.Println("length of arr:", len(arr))

	//
	// Recursion
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var a [5]int
	fmt.Println("emp:", a)

	//
	// Set value of array index
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//
	// Declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	//
	// Multi-dimensional data structures
	var twoD [2][3]int
	fmt.Println("multiD:", twoD) // [[0 1 2] [1 2 3]]

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	//
	// Unfixed array length [...]
	arrWithUnfixedLength := [...]int{1, 2, 3}
	fmt.Println(len(arrWithUnfixedLength))

	//
	// Multi-dimensional array initialization
	// skeleton: [lenA][lenB]type{[lenB]type{value1, value2, ...}, [lenB]type{value1, value2, ...}}
	mdArr := [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	fmt.Println(mdArr)
}
