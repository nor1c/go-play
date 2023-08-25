package main

import (
	"log"
	"time"

)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for c := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				log.Println(i)
			}
		}(c+1, chans[c])
	}

	for {
		select {
		case m0 := <-chans[0]:
			log.Println("received", m0)
		case m1 := <-chans[1]:
			log.Println("received", m1)
		case <-time.After(5 * time.Second):
			log.Println("finished")
			return
		}
	}
}
