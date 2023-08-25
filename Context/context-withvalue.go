package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.WithValue(context.Background(), "language", "Go")
	fmt.Println(ctx1.Value("language"))

	// example 2
	f := func(ctx context.Context, k string) {
		fmt.Println(ctx.Value(k))
	}

	ctx2 := context.WithValue(context.Background(), "language", "Go")

	f(ctx2, "language")

	// example 3
	type ctxKey string

	ff := func(ctx context.Context, k ctxKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println(v)
			return
		}

		fmt.Println("no key found!")
	}

	newKey := ctxKey("language")
	ctx3 := context.WithValue(context.Background(), newKey, "Go")

	ff(ctx3, newKey)
}
