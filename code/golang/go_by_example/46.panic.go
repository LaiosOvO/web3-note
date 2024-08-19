package main

import "os"

func main() {

	panic("a problem here ")

	_, error := os.Create("/tmp/file")
	if error != nil {
		panic(error)
	}
}
