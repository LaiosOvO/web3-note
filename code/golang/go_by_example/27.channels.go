package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	//msg := <-messages
	//fmt.Println(msg)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
	messages <- "pong"
	time.Sleep(time.Second)

}
