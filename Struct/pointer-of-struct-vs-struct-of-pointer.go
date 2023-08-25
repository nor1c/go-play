package main

import "fmt"

type Owner struct {
	Name string
}

func main() {
	// []*
	var owners []*Owner

	owners = append(owners, &Owner{Name: "Alice"})
	owners = append(owners, &Owner{Name: "Bob"})

	for _, owner := range owners {
		fmt.Println(owner.Name)
	}
	// ✅ ok!
	// Alice
	// Bob

	// *[]
	var ows *[]Owner

	// ❕❗❕❗ read this
	// ows = append(ows, &Owner{Name: "Alice"}) // 💥 cannot range over ows (variable of type *[]Owner)

	owsSlice := []Owner{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	ows = &owsSlice

	for _, owner := range *ows {
		fmt.Println(owner.Name)
	}
}
