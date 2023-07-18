package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1 // print 1-3
	}

	for j := 1; j <= 5; j++ {
		fmt.Println(j) // print 1-5
	}

	for {
		fmt.Println("loop")
		break // will repeatedly loop until you break it out
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n) // 1, 3, 5
	}

	// init and post statement are optional
	ii := 1
	for ii < 1000 { // same as `for ; ii < 1000;`
		ii += ii
	}
	fmt.Println("ii:", ii) // 1024

	//
	// Outerloop (exit or stop the loop on nested loop)
outerLoop:
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if i == 2 && j == 2 {
				break outerLoop
			}
			fmt.Println(i, j)
		}
	}
	/*
		output:
		1 1
		1 2
		1 3
		1 4
		2 1
		// stopped when i is 2 and j is 2, so it won't print the 2 2
	*/
}
