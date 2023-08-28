package main

import (
	"fmt"
	"log"
	"os"

)

type Student struct {
	ID    int
	Name  string
	Score float64
}

func main() {
	// currDir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal("err getwd")
	// }

	// fPath := filepath.Join(currDir, "Fmt/students.txt")
	// fmt.Println(fPath)

	//
	file, err := os.Open("Fmt/students.txt")
	if err != nil {
		log.Fatalf("Error opening file. Err: %v", err)
	}
	defer file.Close()

	var students []Student
	for {
		var student Student

		_, err := fmt.Fscanf(file, "%d %s %f", &student.ID, &student.Name, &student.Score)
		if err != nil {
			break
		}

		students = append(students, student)
	}

	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Score: %f\n", student.ID, student.Name, student.Score)
	}
}
