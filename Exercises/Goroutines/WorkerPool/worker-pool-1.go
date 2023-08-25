package main

import (
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func fetchStatus(sites <-chan Site, r chan<- Result) {
	for site := range sites {
		res, err := http.Get(site.URL)
		if err != nil {
			log.Printf("err for %s: %s", site.URL, err.Error())
		}

		r <- Result{
			URL:    site.URL,
			Status: res.StatusCode,
		}
	}
}

func main() {
	siteCh := make(chan Site, 4)
	resultCh := make(chan Result, 4)

	for i := 0; i < 10; i++ {
		go fetchStatus(siteCh, resultCh)
	}

	sitesToCheck := [4]string{
		"https://google.com",
		"https://youtube.com",
		"https://yahoo.com",
		"https://bing.com",
	}

	for _, site := range sitesToCheck {
		siteCh <- Site{URL: site}
	}
	close(siteCh)

	for r := range resultCh {
		log.Printf("status for %s: %d", r.URL, r.Status)
	}
}
