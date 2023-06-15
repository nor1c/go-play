package main

import "fmt"

func main() {
	defer fmt.Println("[deferred] world")
	defer fmt.Println("[deferred]", 1)
	defer fmt.Println("[deferred]", 2)

	fmt.Println("Hello")
	// output: Hello 2 1 world (LiFo: Last-in, First-out)

	myDefer() // 43210
	fmt.Println()

	fmt.Println("Opening file 1")
	defer fmt.Println("[deferred] Closing file 1")

	fmt.Println("Opening file 2")
	defer fmt.Println("[deferred] Closing file 2")

	deferI()

	fmt.Println(c())
}

func deferI() {
	i := 0
	defer fmt.Println(i) // 0
	i++
	// fmt.Println(i) // if we place here, it will return 1, because it's already incremented by 1 (i++)
	return
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}
