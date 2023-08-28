package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

)

var (
	name string = "fauzi"
	age  int    = 25
)

func main() {
	// Errorf: formatting to create descriptive error messages.
	e1 := fmt.Errorf("User %d is not found!", 2)
	fmt.Println(e1.Error()) // User 2 is not found

	// difference with errors.New
	en := errors.New("Some error")
	fmt.Println(en.Error())

	fmt.Println()

	// return bytes/length of a value, any additional character like `\n` still counted as 1 character
	// this doesn't support formatted string unlike Fprintf
	n1, err := fmt.Fprint(os.Stdout, "test")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Fprint: %v\n", err)
	}
	fmt.Printf("\nBytes of n1 is: %v\n", n1) // 4
	fmt.Println()

	//
	fmt.Fprintf(os.Stdout, "%s is %d years olds", name, age) // fauzi is 25 years old

	// print value and return bytes/length
	n2, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old")
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
	}
	fmt.Printf("Bytes of n2 is: %v\n", n2) // 22
	fmt.Println()

	// Fscanf (scan from reader)
	var (
		ID    int
		Name  string
		Score float64
	)
	r := strings.NewReader("103 AhmadFauzi 95.5")
	n, err := fmt.Fscanf(r, "%d %s %f", &ID, &Name, &Score)
	if err != nil {
		log.Fatalf("Fscanf err: %v", err)
	}
	fmt.Println(ID, Name, Score) // 103 AhmadFauzi 95.5
	fmt.Println(n)               // 3
	fmt.Println()

	// Sprint
	s := fmt.Sprint("My name is ", name)
	fmt.Println(s) // My name is fauzi

	// Sprintf
	sf := fmt.Sprintf("My name is %v", name)
	fmt.Println(sf) // My name is fauzi

	// Sprintln
	sl := fmt.Sprintln("My name is ", name)
	fmt.Println(sl) // My name is fauzi (+with line break)

	// 
}
