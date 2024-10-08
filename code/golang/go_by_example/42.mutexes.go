package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func main() {

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)

	doIncrement("a", 100_0000)
	doIncrement("b", 100_0000)
	doIncrement("a", 100_0000)

	wg.Wait()
	fmt.Println("all done\n", c.counters)
}
