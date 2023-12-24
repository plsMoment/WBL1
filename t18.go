package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Counter struct like lazy atomic only with increment:)
type Counter struct {
	count int
	m     sync.Mutex
}

func (c *Counter) Inc() {
	c.m.Lock()
	c.count++
	c.m.Unlock()
}

func (c *Counter) String() string {
	return strconv.Itoa(c.count)
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c)
}
