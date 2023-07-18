package main

import "fmt"

func main() {
	// zero value of slices is "nil"
	var anEmptySlices []int
	fmt.Println(anEmptySlices, len(anEmptySlices), cap(anEmptySlices)) // [] 0 0

	s := make([]string, 3)
	fmt.Println("emp:", s) // [   ]

	n := make([]int, 3)
	fmt.Println("emp:", n) // [0 0 0]

	// len and cap
	lc := make([]int, 0, 5)
	fmt.Println(lc, len(lc), cap(lc)) // [] 0 5

	lc2 := make([]int, 3, 5)
	fmt.Println(lc2, len(lc2), cap(lc2)) // [0 0 0] 3 5

	lc3 := make([]int, 5)
	fmt.Println(lc3, len(lc3), cap(lc3)) // [0 0 0 0 0] 5 5 (cap automatically set to 5, based on the len of the slice)

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

	sc = append(sc, 5, 6) // current capacity: 6
	sc = append(sc, 7)    // current capacity: 12. No more space, so it will expand its size, and doubled from last capacity size which is 6

	//
	// Copying a slice (copy())
	/*
		the skeleton:
		sliceToCopy := []type{value1, value2, value3}
		newSliceTemplate := make([]type, length)
		copy(newSliceTemplate, sliceToCopy)
	*/
	c := make([]string, len(s)) // empty slice with the same length as 's'
	copy(c, s)
	fmt.Println("copy of s:", c) // [first second third fourth fifth sixth]

	sliceToCopy := []int{1, 2, 3}
	destSlice := make([]int, 2)
	copy(destSlice, sliceToCopy)
	fmt.Println(destSlice) // [1 2]

	// Slice copy & mutating
	mutSource := []int{1, 2}
	mutDst := []int{4, 5, 6}
	copy(mutDst, mutSource)
	fmt.Println(mutDst) // 1, 2, 6

	/*
		slice operators
	*/
	sl1 := s[2:5]            // will take the value of slice 's' from index 2 to index 4 (2 -> <= 5) (s[2], s[3], and s[4])
	fmt.Println("sl1:", sl1) // [third fourth fifth]

	sl2 := s[:5]             // index 0 to 4 (s[0], s[1], s[2], s[3], s[4])
	fmt.Println("sl2:", sl2) // [first second third fourth fifth]

	sl3 := s[2:]             // index 3 to last index
	fmt.Println("sl3:", sl3) // [third fourth fifth sixth]

	sl4 := s[:]                            // all elements
	fmt.Println("sl4 (all elements)", sl4) // [first second third fourth firth sixth]

	// slice operator with cap
	sl5 := s[0:1:2]
	fmt.Println("sl5:", sl5)                  // [first]
	fmt.Println("capacity of sl5:", cap(sl5)) // 2

	/*
		declare and initialize a slice in single line of code
	*/
	t := []string{"g", "h", "i"}
	fmt.Println(t) // [g h i]

	// slices is like a references, when you change its element it's also change the original slices elements
	// not like array that created a new copy of array
	ogSlices := []int{1, 2, 3}
	fmt.Println(ogSlices) // [1 2 3]

	ogSlices[1] = 4
	fmt.Println(ogSlices) // [1 4 3]

	// multi-dimensional data structure
	twoDSlices := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoDSlices[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoDSlices[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoDSlices) // [[0] [1 2] [2 3 4]]

	// a slice of anonymous struct
	structSlices := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}

	fmt.Println(structSlices) // [{2 true} {3 false}]

	// slicing a slice
	fullSlice := []int{1, 2, 3, 4, 5}

	fullSlice = fullSlice[:0]
	fmt.Println(fullSlice) // []

	/**
	fullSlice[1] = 1 // error: because len of fullSize is 0
	*/

	fullSlice = fullSlice[:5]
	fmt.Println(fullSlice) // [1 2 3 4 5]

	fullSlice[1] = 1
	fmt.Println(fullSlice) // [1 1 3 4 5]

	// slices of slices
	board := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board[0]) // [1 2 3]

	board[0][0] = 2
	fmt.Println(board[0]) // [2 2 3]

	board[1][2] = 10
	fmt.Println(board) // [[2 2 3] [4 5 10] [7 8 9]]

	// append slices to another slices
	firstSlices := []int{1, 2, 3}
	secondSlices := []int{4, 5, 6}

	firstSlices = append(firstSlices, secondSlices...)

	fmt.Println(firstSlices) // [1, 2, 3, 4, 5, 6]
}
