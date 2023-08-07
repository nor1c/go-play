package main

import (
	"fmt"
	"sync"
	"time"
)

func runWorker(wg *sync.WaitGroup, n int) {
	wg.Add(1)
	defer wg.Done()
	fmt.Printf("Worker %v running..\n", n)
	time.Sleep(time.Second * 2)
	fmt.Printf("Worker %v done!\n", n)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		go runWorker(&wg, i)
	}

	wg.Wait()
}
