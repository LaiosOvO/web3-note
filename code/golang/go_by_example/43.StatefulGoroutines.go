package main

import (
	"fmt"
	"math/rand/v2"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	/**
	a goroutine own the state variable
	keep reading and writing data
	*/
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// read gorountine
	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := readOp{
					key:  rand.IntN(5),
					resp: make(chan int),
				}
				reads <- read
				fmt.Println("received ", <-read.resp)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// write gorountine
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.IntN(5),
					val:  rand.IntN(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// main goroutine
	time.Sleep(3 * time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)

}
