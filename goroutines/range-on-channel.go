package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	user := make(chan string, 5)

	wg.Add(3)
	go getUser(user, &wg)
	go getRandomuser(user, &wg)
	go getTopUser(user, &wg)

	wg.Wait()
	close(user)

	for u := range user {
		fmt.Println(u)
	} // output: dwayne josh fauzi
}

func getUser(user chan string, wg *sync.WaitGroup) {
	user <- "dwayne"
	wg.Done()
}

func getRandomuser(user chan string, wg *sync.WaitGroup) {
	user <- "josh"
	wg.Done()
}

func getTopUser(user chan string, wg *sync.WaitGroup) {
	user <- "fauzi"
	wg.Done()
}
