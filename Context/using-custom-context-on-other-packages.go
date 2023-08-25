package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://example.com/", nil)
	if err != nil {
		log.Fatalf("err: %v", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// we can apply our `ctx` to `req` using `.WithContext()`
	// most packages has this method
	req = req.WithContext(ctx)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("err (c.Do):", err.Error())
	}
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("err (out):", err.Error())
	}

	fmt.Println(string(out))
}
