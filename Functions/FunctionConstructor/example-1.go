package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
		Age:  0,
	}
}

func NewStudentWithAge(name string, age int) *Student {
	return &Student{
		Name: name,
		Age:  age,
	}
}

func main() {
	fauzi := NewStudent("A. Fauzi")
	fmt.Println("Student name:", fauzi.Name, "age:", fauzi.Age)

	jack := NewStudentWithAge("Jack Kambey", 48)
	fmt.Println("Student name:", jack.Name, "Age:", jack.Age)
}
