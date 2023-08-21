package main

import "fmt"

func sumNonGeneric(numbers []int) int {
	inc := 0
	for _, v := range numbers {
		inc += v
	}
	return inc
}

func sumGeneric[T int](numbers []T) T {
	var inc T
	for _, v := range numbers {
		inc += v
	}
	return inc
}

func main() {
	fmt.Printf("%d\n", sumNonGeneric([]int{1, 2, 3, 4}))
	fmt.Printf("%d", sumGeneric([]int{1, 2, 3, 4}))
}
