package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int32
}

func main() {
	var emptyInterface interface{}
	fmt.Println(emptyInterface) // default value is nil
	fmt.Printf("Type of 'emptyInterface': %T\n", emptyInterface)

	// interface{} is the same as any
	var interfaceIsAny any = "anything you like"
	fmt.Println(interfaceIsAny)
	fmt.Printf("Type of 'interfaceIsAny': %T", interfaceIsAny)

	// interface{} can accept any value (string, number, pointer, data object, etc.)
	var i1 interface{}
	i1 = []string{"apple", "mango", "banana"}
	fmt.Println(i1)

	// we need to assert the type of the interface{} variable first before doing some operation on it
	// we can assert type like this: `variable.(type)`
	fmt.Println(strings.Join(i1.([]string), " ")) // output: apple mango banana

	var mike interface{} = &Person{
		Name: "Mike",
		Age:  32,
	}
	// fmt.Println(mike.Name) // error
	fmt.Println(mike.(*Person).Name) // output: Mike
}
