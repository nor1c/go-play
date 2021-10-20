package main

import (
	"fmt"
	"reflect"
)

func main() {
	var singleVariableNoType = "golang"
	fmt.Println(singleVariableNoType, reflect.TypeOf(singleVariableNoType)) // golang string

	var multiVarDeclaration1, multiVarDeclaration2 = 1, "golang"
	fmt.Println(reflect.TypeOf(multiVarDeclaration1), reflect.TypeOf(multiVarDeclaration2)) // int string
	fmt.Println(multiVarDeclaration1, multiVarDeclaration2)                                 // 1 golang

	var intA, intB int = 1, 2
	fmt.Println(reflect.TypeOf(intA), reflect.TypeOf(intB)) // int int
	fmt.Println(intA, intB)                                 // 1, 2
	intA = 3
	fmt.Println(intA) // 3

	var booleanVariable = true
	fmt.Println(booleanVariable, reflect.TypeOf(booleanVariable)) // true bool

	var undefinedVariable int
	fmt.Println(undefinedVariable)

	shorthand := "variable declaration using shorthand"
	fmt.Println(shorthand)
	shorthand = "is it changed?"
	if shorthand == "is it changed?" {
		fmt.Println("is it changed? yes")
	} else {
		fmt.Println("is it changed? no")
	} // is it changed? yes
}
