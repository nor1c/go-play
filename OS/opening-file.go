package main

import (
	"log"
	"os"

)

func main() {
	file, err := os.Open("filetoopen.txt")
	if err != nil {
		log.Fatalf("Failed to open file. Err: %v", err)
	}
	defer file.Close()
}
