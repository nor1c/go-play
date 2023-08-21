package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

End:
	fmt.Println("is done", <-done)

	for i := 1; i <= 5; i++ {
		if i == 4 {
			go func() {
				done <- true
			}()
			goto End
		} else {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}
}
