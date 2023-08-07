package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	fmt.Println("⚪ Executing Goroutine")
	fmt.Println("🟡 Delaying for 2 seconds (simulating expensive processes)")
	time.Sleep(time.Second * 2)
	fmt.Println("🟢 Finished executing Goroutine")

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go myFunc(&wg)

	wg.Wait()
	fmt.Println("🟢 Finished executing my go program!")
}
