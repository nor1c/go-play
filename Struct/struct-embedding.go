package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("desribed: num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 2,
		},
		str: "Hello, World!",
	}
	fmt.Printf("co={num=%v, str=%v}\n", co.base.num, co.str)
	fmt.Printf("also num=%v\n", co.num) // 2

	fmt.Println(co.describe()) // desribed: num=2

	// base
	b := base{
		num: 3,
	}
	fmt.Println(b.describe()) //described: num=3

	// interface implementations onto other struct
	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println(d.describe()) // described: num=2
}
