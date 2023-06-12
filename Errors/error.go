package main

import (
	"errors"
	"fmt"
	"log"
)

var notOddNumberError = errors.New("The number you provided is not an odd number")

func main() {
	// New
	err1 := errors.New("Something went wrong!")

	if err1 != nil {
		fmt.Println(err1)
	}

	//
	divided, errDivided := Divide(10, 2)
	if errDivided != nil {
		log.Fatal(errDivided)
	}
	fmt.Println(divided)

	// check is the error is `notOddNumberError`
	isOdd, oddErr := IsOdd(5)
	if oddErr != nil {
		if errors.Is(oddErr, notOddNumberError) {
			log.Fatal(notOddNumberError)
		} else {
			log.Fatal("ODD: Unknown error", oddErr)
		}
	}

	fmt.Printf("is odd: %t", isOdd)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide %d by zero", b)
	}

	return a / b, nil
}

func IsOdd(number int) (bool, error) {
	if number%2 == 0 {
		return true, notOddNumberError
	}

	return true, nil
}
