package main

import (
	"context"
	"fmt"

)

func enrichContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "request-id", "12345")
	ctx = context.WithValue(ctx, "name", "Fauzi")

	return ctx
}

func main() {
	ctx := context.Background()

	ctx = enrichContext(ctx)

	fmt.Println(ctx.Value("request-id"))
	fmt.Println(ctx.Value("name"))
}
