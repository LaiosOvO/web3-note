package main

import (
	"fmt"
)

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) getAll() []T {
	var elems []T

	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys: ", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.push(10)
	lst.push(20)
	lst.push(30)
	lst.push(40)
	fmt.Println(lst.getAll())

}

func m1(a int) int {
	return 1
}
