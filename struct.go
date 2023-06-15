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
}
