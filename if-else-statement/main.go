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
}
