package main

import "fmt"

func valus() (int, int) {
	return 3, 7
}

func main() {

	a, b := valus()
	fmt.Println(a)
	fmt.Println(b)

	_, c := valus()
	fmt.Println(c)

}
