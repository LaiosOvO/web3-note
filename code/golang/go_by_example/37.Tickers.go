package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// 创建一个Ticker，每秒钟触发一次
	//ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at ", t)
			}
		}
	}()

	time.Sleep(1500 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")

}
