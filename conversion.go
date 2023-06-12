package main

import "fmt"

func main() {
	/**
	%v - any format
	%d - number/decimal/int
	%g - floating number
	%b - base 2 numbers
	%o - base 8 numbers
	%t - bool
	%s - string
	*/
	vNum, vTrue, vFalse, vString := 2, true, false, "welcome!"
	fmt.Println(fmt.Sprintf("Number: %d, Boolean: %t & %t, String: %s", vNum, vTrue, vFalse, vString))

	// %v can accept anything
	fmt.Println(fmt.Sprintf("Number: %v, Boolean: %v & %v, String: %v", vNum, vTrue, vFalse, vString))

	pi := 3.14
	fmt.Println(fmt.Sprintf("Pi: %g", pi))

	valueNumber, valueString := 2, "a string"
	fmt.Println(fmt.Sprintf("%d, %s", valueNumber, valueString))
}
