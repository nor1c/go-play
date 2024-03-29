package main

import (
	"fmt"
	"io"
	"net/http"

)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}
