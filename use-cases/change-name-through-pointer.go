package main

import "fmt"

func main() {
	name := "fauzi"

	fmt.Println(name)

	changeName(&name)

	fmt.Println(name)
}

func changeName(name *string) {
	*name = "A. Fauzi"
}
