package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	})

	srv := &http.Server{
		Addr: ":8080",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
