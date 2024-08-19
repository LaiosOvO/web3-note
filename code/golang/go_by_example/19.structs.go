package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 18
	return &p
}

func main() {

	fmt.Println(person{"laios1", 12})
	fmt.Println(person{name: "laios1", age: 12})
	fmt.Println(person{name: "laios1"})
	fmt.Println(&person{name: "laios1", age: 22})
	fmt.Println(newPerson("laioss"))

	s := person{name: "laios", age: 12}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 22
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"tom",
		true,
	}
	fmt.Println(dog)

}
