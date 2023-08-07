package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://facebook.com",
	"https://youtube.com",
}

func fetch(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%v -- %+v\n", url, resp.Status)
		}(url)
	}

	wg.Wait()
}

func main() {
	http.HandleFunc("/", fetch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
