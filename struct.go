package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	anita := Person{"Anita", 20}
	fmt.Println(anita) // {Bob 20}

	// Anita married with William, so her new name is Anita William
	anitaWilliam := &anita
	anitaWilliam.Name = "Anita Willian"

	fmt.Println(anita.Name)        // Anita William
	fmt.Println(anitaWilliam.Name) // Anita William

	fmt.Println(reflect.TypeOf(anita))        // main.Person
	fmt.Println(reflect.TypeOf(&anita))       // *main.Person
	fmt.Println(reflect.TypeOf(anitaWilliam)) // *main.Person

	// anonymous struct
	var dog struct {
		Legs   int
		CanRun bool
	}
	dog.Legs = 4
	dog.CanRun = true
	fmt.Println(dog) // {4 true}

	// slice bounds
	sliceBounds := []int{2, 3, 5, 6, 7}

	sliceBounds = sliceBounds[1:4]
	fmt.Println(sliceBounds) // [3 5 6 7]

	sliceBounds = sliceBounds[:2]
	fmt.Println(sliceBounds) // [3 5]

	sliceBounds = sliceBounds[1:]
	fmt.Println(sliceBounds) // [5 6]
}
