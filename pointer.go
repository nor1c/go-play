package main

import "fmt"

func main() {
	//
	//
	// example 1
	i, j := 42, 1000

	p := &i         // 0xc00001e098 (pointer)
	fmt.Println(*p) // read i through pointer

	*p = 21            // set i through pointer
	fmt.Println(p, *p) // 0xc0000a...., 21
	fmt.Println(i)     // new i will be 21

	pJ := &j
	fmt.Println(*pJ) // 1000

	*pJ = *pJ / 10
	fmt.Println(*pJ) // 100
	fmt.Println(j)   // 100

	//
	//
	// example 2
	var xx int = 42
	fmt.Println(xx)  // 42
	fmt.Println(&xx) // memory address of `xx`

	var pp *int     // declare a pointer variable that can store the memory address of an integer variable
	fmt.Println(pp) // nil

	pp = &xx         // assign the memory address of `xx` to pointer `pp`
	fmt.Println(pp)  // memory address of `xx`
	fmt.Println(&pp) // memory address of `pp`
	fmt.Println(*pp) // 42 (value of `xx`)
	/**
	xx = value
	pp = nil
	&pp = memory address
	*pp = value
	*/
}
