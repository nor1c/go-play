package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// basic example of how to use switch in go
	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// multiple expressions in the same case statement
	today := time.Now().Weekday()
	fmt.Println(today)

	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch with variable declaration
	switch os := runtime.GOOS; os {
	case "darwing":
		fmt.Println("OS X")
	default:
		fmt.Println("Linux/Windows")
	}

	// switch without an expression
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(2)
	whatAmI(true)
	whatAmI("hey")

	//
	// Multiple condition in switch case
	point := 8
	switch point {
	case 6:
		fmt.Println("The point is 6")
	case 7, 8:
		fmt.Println("The point is 7 or 8")
	default:
		fmt.Println("The point isn't 6, 7 or 8")
	} // The point isn't 6, 7 or 8

	//
	// Case with if-else condition style
	switch { // you don't need to define expression here
	case point <= 5:
		fmt.Println("The point is less than or equal to 5")
	case point >= 6 && point <= 10:
		fmt.Println("The point is bigger than 5 and less than or equal to 10")
	default:
		fmt.Println("The point is bigger than 10")
	} // The point is bigger than 5 and less than or equal to 10

	//
	// Fallthrough (force to run next case code)
	switch {
	case point <= 5:
		fmt.Println("Point is less than or equal to 5")
	case point >= 6 && point <= 10:
		fmt.Println("Point is bigger than 6 and less than or equal to 10")
		fallthrough // this will force to run next code no matter what
	case point == 5:
		fmt.Println("[FORCED] Point is 5")
	case point == 3:
		fmt.Println("Point is 3")
	default:
		fmt.Println("Point is bigger than 10")
	}
	/*
		output:
		Point is bigger than 6 and less than or equal to 10
		[FORCED] Point is 5
	*/
}
