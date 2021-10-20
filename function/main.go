package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

// variadic function
func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	res := plus(1, 2)
	fmt.Println(res) // 3

	res3 := plusPlus(1, 2, 3)
	fmt.Println(res3) // 6

	// return multiple values
	first, second := vals()
	fmt.Println(first, second) // 3, 7

	// return specific index only
	_, secondsecond := vals()
	fmt.Println(secondsecond) // 7

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...) // [1 2 3 4 5] 15
	// or
	sum(1, 2, 3, 4) // [1 2 3 4] 10
}
