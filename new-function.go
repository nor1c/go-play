package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func main() {
	// new(), create a new instance and return a pointer to that value
	// new(), allocate memory and intialize it to the zero value of that type (0 for int, and "" for string)

	// example:
	var person Person
	p := &person
	fmt.Println(p) // &{"" 0}

	person2 := &Person{}
	fmt.Println(person2) // &{"" 0}

	person3 := new(Person)
	fmt.Println(person3) // &{"" 0}
}
