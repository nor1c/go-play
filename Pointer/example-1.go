package main

import "fmt"

func changeName(name *string) {
	fmt.Println("mem(f):", &name) // it will return different mem address as line 21, but it's still act as a reference
	*name = "a. fauzi"
}

func changeNameNonPointer(name string) {
	name = "fauzi (non-p)"
	fmt.Println("name changed to:", name)
}

func main() {
	var name string

	name = "fauzi"

	fmt.Println(name)
	fmt.Println("mem:", &name) // return original mem address of var `name`

	changeName(&name)              // name changed cuz it passes a reference of variable name (it's not copied the variable/recreating a new variable)
	fmt.Println("new name:", name) // 'a. fauzi'
	fmt.Println("mem:", &name)     // return same mem address as line 21

	changeNameNonPointer(name)             // won't change the name cuz it's copied the variable (not as a reference)
	fmt.Println("new name (non-p):", name) // still 'a. fauzi'
}
