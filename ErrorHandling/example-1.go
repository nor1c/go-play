package main

import (
	"fmt"
)

func SomeTask(n1 int, n2 int) (int, error) {
	if n1 > 1 {
		return 0, fmt.Errorf("n1 can't bigger than 1, received %d", n1)
	} else {
		return n1 + n2, nil
	}
}

func main() {
	n1 := 1
	n2 := 2
	sum1, err := SomeTask(n1, n2)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("result: %d", sum1)
}
