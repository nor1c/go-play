package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Println(err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(string(body))
}