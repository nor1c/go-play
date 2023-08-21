package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func NewPerson(name string, age uint8) *Person {
	person := &Person{
		Name: name,
		Age:  age,
	}
	fmt.Println("mem of NewPerson (person1):", &person) // 0x000...xx

	return person
}
func changeName(person *Person, newName string) {
	fmt.Println("mem changeName()", &person.Name) // same as the original mem address of person1 in the main func
	person.Name = newName
}

func NewPersonNonPointer(name string, age uint8) Person {
	person := Person{
		Name: name,
		Age:  age,
	}
	fmt.Println("mem of NewPersonNonPointer (person2):", &person) // &{jack 45}

	return person
}
func changeNameNonPointer(person Person, newName string) {
	fmt.Println("mem changeNameNonPointer()", &person.Name) // different mem address (new copy)
	person.Name = newName
}

func main() {
	person1 := NewPerson("a. fauzi", 26)
	fmt.Println("mem of person1:", &person1) // returning a memory address 0xc000...xx
	fmt.Println("mem person1.Name", &person1.Name)
	fmt.Println(person1.Name) // 'a. fauzi'
	changeName(person1, "fauzi")
	fmt.Println("new name:", person1.Name) // 'fauzi' ✅

	fmt.Println("")

	person2 := NewPersonNonPointer("jack", 45)
	fmt.Println("mem of person2:", &person2) // &{jack 45} <- it's not returning a memory address
	fmt.Println("mem person2.Name", &person2.Name)
	fmt.Println(person2.Name) // 'jack'
	changeNameNonPointer(person2, "jack harrington")
	fmt.Println("new name:", person2.Name) // 'jack' ❌ it's still 'jack' because in the `changeNameNonPointer` they created a new copy of the variable passed instead of the reference
}
