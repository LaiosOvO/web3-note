package main

import "fmt"

/**
闭包可以捕获其创建我在作用域的变量，及时这些变量在闭包创建后已经离开了作用域
*/

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Print(nextInt())
	fmt.Print(nextInt())
	fmt.Print(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())

}
