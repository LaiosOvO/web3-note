package main

import "fmt"

type base struct {
	num int
}

func (b base) describte() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Println(co.num, co.str)
	fmt.Println(co.base.num)

	fmt.Println(co.describte())

	type describer interface {
		describte() string
	}

	var d describer = co
	fmt.Println(d.describte())
}
