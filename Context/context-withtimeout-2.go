package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingCool(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go doSomethingCool(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("application timed out")
	}
}
