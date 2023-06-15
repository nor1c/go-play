package main

import "fmt"

func main() {
	if 7%2 == 0 { // if else example
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // if statement without else
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // var declaration within if-else statement
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	} // 9 has 1 digit

	// variable declaration in a if statement
	if a := 5; a < 6 {
		fmt.Printf("a (%d) is less than 6\n", a)
	} else {
		fmt.Printf("a (%d) is bigger than 6\n", a)
	}

	fmt.Println("Is 5 even?", isPositiveEven(5))
	fmt.Println("Is 6 even?", isPositiveEven(6))
}

// example of variable declaration in a statement
func isPositiveEven(n int) bool {
	if isPositive, isEven := n > 0, n%2 == 0; isPositive && isEven {
		return true
	}
	return false
}
