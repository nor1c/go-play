package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	aNumber, aString, aBool := 24, "Fauzi", true

	message := "Hello Fauzi! 你好!"

	p := Person{
		Name: "Fauzi",
		Age:  26,
	}

	//
	fmt.Printf("aNumber: %d, aString: %s, aBool: %t\n", aNumber, aString, aBool)

	// print type of the variable
	fmt.Printf("Type of aNumber: %T, aString: %T\n", aNumber, aString)

	// can accept anything
	fmt.Printf("aNumber: %v, aString: %v, aBool: %v\n", aNumber, aString, aBool)

	// v
	fmt.Printf("%+v\n", p) // {Name:Fauzi Age:26}
	fmt.Printf("%#v\n", p) // main.Person{Name:"Fauzi", Age:26}

	// quoted
	fmt.Printf("Message: %q\n", message)  // Message: "Hello Fauzi! 你好!"
	fmt.Printf("Message: %+q\n", message) // Message: "Hello Fauzi! \u4f60\u597d!"
	fmt.Printf("Message: %#q\n", message) // Message: `Hello Fauzi! 你好!`

	// decimal integer
	dec := 42
	fmt.Printf("Decimal: %d\n", dec) // Decimal: 42
	fmt.Printf("Binary: %b\n", dec)  // Binary: 101010
	fmt.Printf("Octal: %o\n", dec)   // Octal: 52
}
