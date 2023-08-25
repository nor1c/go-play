package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func midd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("start: executing middleware..")

		ctx := context.WithValue(context.Background(), "request-id", "123456")

		next.ServeHTTP(w, r.WithContext(ctx))

		log.Println("end: done processing middleware!")
	})
}

func main() {
	mux := http.NewServeMux()

	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rID := r.Context().Value("request-id")
		fmt.Println(rID)

		fmt.Fprintf(w, "Hello, world!")
	})
	mux.Handle("/", midd(rootHandler))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
