package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1 // print 1-3
	}

	for j := 1; j <= 5; j++ {
		fmt.Println(j) // print 1-5
	}

	for {
		fmt.Println("loop")
		break // will repeatedly loop until you break it out
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n) // 1, 3, 5
	}
}
