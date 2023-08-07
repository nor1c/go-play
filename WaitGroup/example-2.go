package main

import (
	"fmt"
	"sync"
	"time"
)

func printS(s string) {
	fmt.Printf("%v \n", s)
	time.Sleep(time.Second * 2)
	fmt.Printf("%v done!\n", s)
}

func main() {
	var wg sync.WaitGroup

	ss := []string{"a", "b", "c"}
	for _, s := range ss {
		wg.Add(1)

		go func(s string) {
			defer wg.Done()
			printS(s)
		}(s)
	}

	wg.Wait()
	fmt.Println("🟢 Finished running my Go program!")
}
