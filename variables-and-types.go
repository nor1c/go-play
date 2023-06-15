package main

import (
	"fmt"
	"reflect"
)

func main() {
	// default value of variables with type
	var (
		ds string
		dn int
		db bool
		df float64
	)
	fmt.Printf("%s %d %t %f\n", ds, dn, db, df) // "" 0 false 0

	// type on multiple variable as one
	var canTalk, canSwim bool
	canTalk, canSwim = false, true
	fmt.Println(canTalk, canSwim) // false true

	// multiple variable but with different types
	var (
		leg        uint
		canRun     bool
		height     float64
		animalName string
	)
	leg, canRun, height, animalName = 4, true, 1.63, "Horse"
	fmt.Printf("%s has %d legs, height: %gm, and can run? %t\n", animalName, leg, height, canRun) // Horse has 4 legs, height: 1.63m, and can run? true

	// multiple variables with initializers
	var (
		mA bool   = true
		mB string = "Hello"
		mC int    = 40
	)
	fmt.Println(mA, mB, mC) // true, "Hello", 40

	// variables with initializers
	var firstNumber, secondNumber uint = 1, 2
	fmt.Printf("First number: %d, Second number: %d\n", firstNumber, secondNumber) // First number: 1, Second number: 2

	// short assignment (:=), note that this doesn't work outside func
	shortAssignedString := "Hello, world!"
	fmt.Println(shortAssignedString) // Hello, world!

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
	fmt.Println(undefinedVariable) // 0

	shorthand := "variable declaration using shorthand"
	fmt.Println(shorthand)
	shorthand = "is it changed?"
	if shorthand == "is it changed?" {
		fmt.Println("is it changed? yes")
	} else {
		fmt.Println("is it changed? no")
	} // is it changed? yes
}
