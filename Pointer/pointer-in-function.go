package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

// this will make a copy of passed "person" which is "jono" in this case
// any modification inside this function will only affect the local copy of the original variable passed to the function "person"
func gettingOlderWithoutPointer(person Person) {
	person.age += 1
}

// operates directly to the original variable "person" which in this case is "jono"
// any modification will affect the original variable because they are referring to the same memory location
func gettingOlderWithPointerAddress(person *Person) {
	person.age += 1
}

func main() {
	jono := Person{"Jono", 28}

	gettingOlderWithoutPointer(jono)
	fmt.Println("Without pointer:", jono.age) // 28

	gettingOlderWithPointerAddress(&jono)
	fmt.Println("With pointer:", jono.age) // 29
}
