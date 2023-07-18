package main

import "fmt"

type Person struct {
	Name string
	Age  byte
}

func main() {
	m := make(map[string]int) // map[key-type]val-type, m[string] = int

	m["k1"] = 7
	m["k2"] = 8
	fmt.Println("map:", m) // map[k1:7 k2:8]

	// get value of map key
	v1 := m["k1"]
	fmt.Println("value of m['k1']:", v1) // 7

	fmt.Println("len:", len(m)) // 2

	// delete a key-value pair
	delete(m, "k2")
	fmt.Println("m after deleted k2:", m) // map[k1:7]

	k2Value, isK2Exist := m["k2"]
	fmt.Println(k2Value)           // 0
	fmt.Println("prs:", isK2Exist) // false

	k1Value, isK1Exist := m["k1"]
	fmt.Println(k1Value)            // 7
	fmt.Println("prs1:", isK1Exist) // true

	// declare and initialize map in a single line of code
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n) // map[bar:2 foo:1]

	//
	// Slice of Map
	sliceOfMap := []map[string]string{
		{"name": "Fauzi", "age": "25"},
		{"name": "John Doe", "age": "45"},
	}
	fmt.Println(sliceOfMap) // [map[age:25 name:Fauzi] map[age:45 name:John Doe]]

	for _, k := range sliceOfMap {
		fmt.Println(k["name"], k["age"])
	}
	/**
	output:
	Fauzi 25
	John Doe 45
	*/

	//
	// Map with Struct
	personMap := []map[string]Person{
		{
			"fauzi": Person{Name: "A. Fauzi", Age: 25},
		},
		{
			"john": Person{Name: "John Doe", Age: 46},
		},
	}

	for _, vPerson := range personMap {
		fmt.Println(vPerson)
	} // output: map[fauzi:{A. Fauzi 25} john:{John Doe 46}]

	fmt.Println(personMap[0]["fauzi"]) // {A. Fauzi 25}
}
