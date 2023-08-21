package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		if i == 4 {
			goto End
		} else {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}

End:
	fmt.Println("process ended")
}
