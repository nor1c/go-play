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
}
