package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	Val int
	Err error
}

func main() {
	start := time.Now()

	ctx := context.Background()

	val, err := fetch(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("result:", val)
	log.Println("took:", time.Since(start))
}

func fetch(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	ch := make(chan Response)

	go func() {
		val, err := fetchThatMayTakeTooLong(ch)
		ch <- Response{
			Val: val,
			Err: err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("process timeout because it took to long")
		case r := <-ch:
			return r.Val, r.Err
		}
	}
}

func fetchThatMayTakeTooLong(ch <-chan Response) (int, error) {
	time.Sleep(time.Millisecond * 500) // lower this down <200 to make it return success

	return 333, nil
}
