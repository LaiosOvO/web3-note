package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "b", "a"}
	slices.Sort(strs)
	fmt.Println("string", strs)

	ints := []int{4, 2, 7, 9}
	slices.Sort(ints)
	fmt.Println("ints:  ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("sorted: ", s)
}
