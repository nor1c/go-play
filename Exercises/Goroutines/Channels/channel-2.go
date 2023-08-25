package main

import (
	"fmt"
	"time"
)

func main() {
	materials := [5]string{"rock", "ore", "ore", "rock", "ore"}

	oreCh := make(chan string)
	minedOreCh := make(chan string)

	// finder
	go func(materails [5]string) {
		count := 1
		for _, material := range materials {
			if material == "ore" {
				fmt.Println("🕵️‍♂️ ore found:", count)
				count += 1

				oreCh <- material
			}
		}
	}(materials)

	// ore breaker
	go func() {
		count := 1
		for ore := range oreCh {
			fmt.Println("🔨 ore broken down:", count)
			count += 1

			minedOreCh <- ore
		}
	}()

	// smelter
	go func() {
		count := 1

		for minedOre := range minedOreCh {
			fmt.Printf("🔥 smelted %s: %d\n", minedOre, count)
			count += 1
		}
	}()

	<-time.After(time.Second * 5)
}
