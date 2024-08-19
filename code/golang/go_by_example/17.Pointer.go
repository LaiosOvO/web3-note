package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("initial: ", i)

	zeroVal(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)

	fmt.Println(&i)

}
