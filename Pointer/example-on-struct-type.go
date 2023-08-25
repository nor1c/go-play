package main

import "fmt"

type P struct {
	Name string
}

// copying a new reference
func NewP(name string) P {
	return P{
		Name: name,
	}
}

// using `*P` in this code will not work, since it's not a reference of type `P`
// this will create a new copy of `P`, instead of using the reference
func ChangeP(p P, newName string) {
	p.Name = newName
}

type A struct {
	Name string
}

func NewA(name string) *A {
	return &A{
		Name: name,
	}
}

// using the reference passed
func ChangeA(a *A, newName string) {
	a.Name = newName
}

func main() {
	p := NewP("Fauzi")
	fmt.Println(p.Name) // Fauzi

	ChangeP(p, "Fauzi (new)")
	fmt.Println(p.Name) // Fauzi (❌ still the same, because on ChangeP, it's copying a new reference instead of using current `p`)

	a := NewA("Charlie")
	fmt.Println(a.Name) // Charlie

	ChangeA(a, "Charlie (new)")
	fmt.Println(a.Name) // Charlie (✔ changed)
}
