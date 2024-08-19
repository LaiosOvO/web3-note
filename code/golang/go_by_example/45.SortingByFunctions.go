package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {

	// define a compare function to sort
	fruits := []string{"peach", "banana", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "alios", age: 23},
		Person{name: "blios", age: 100},
		Person{name: "aclios", age: 25},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)

}
