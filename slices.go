package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s) // [   ]

	n := make([]int, 3)
	fmt.Println("emp:", n) // [0 0 0]

	// example in array
	var arrExample [3]int
	fmt.Println("emp arr:", arrExample) // [0 0 0]

	// set and get
	s[0] = "first"
	s[1] = "second"
	s[2] = "third"
	fmt.Println("get second index of s:", s[1]) // second

	// print the length of a slice
	fmt.Println(len(s)) // 3

	// append
	s = append(s, "fourth")
	s = append(s, "fifth", "sixth")
	fmt.Println("apded:", s) // [first second third fourth fifth sixth]

	// append can work in array too
	var apdArr []int
	apdArr = append(apdArr, 1)
	apdArr = append(apdArr, 2, 3)
	fmt.Println("appended array:", apdArr) // [1 2 3]

	// cap() / capacity of slice
	sc := make([]int, 3)                    // base capacity is 3
	fmt.Println("capacity of sc:", cap(sc)) // 3
	// when the slice is expanded it will add more capacity/times based on how many base capacity is
	sc = append(sc, 4)
	fmt.Println("expanded capacity of sc:", cap(sc)) // 6

	// copy a slice
	c := make([]string, len(s)) // empty slice with the same length as 's'
	copy(c, s)
	fmt.Println("copy of s:", c) // [first second third fourth fifth sixth]

	// slice operator
	sl1 := s[2:5]            // will take the value of slice 's' from index 2 to index 4 (2 -> <= 5) (s[2], s[3], and s[4])
	fmt.Println("sl1:", sl1) // [third fourth fifth]

	sl2 := s[:5]             // index 0 to 4 (s[0], s[1], s[2], s[3], s[4])
	fmt.Println("sl2:", sl2) // [first second third fourth fifth]

	sl3 := s[2:]             // index 3 to last index
	fmt.Println("sl3:", sl3) // [third fourth fifth sixth]

	// declare and initialize a slice in single line of code
	t := []string{"g", "h", "i"}
	fmt.Println(t) // [g h i]

	// multi-dimensional data structure
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD) // [[0] [1 2] [2 3 4]]
}
