package main

import (
	"fmt"
	"sync"
)

// mul2 multiply by 2 value from in channel and send it in channel out
func mul2(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val << 1
	}
	close(out)
}

func read(out <-chan int, wg *sync.WaitGroup) {
	for val := range out {
		fmt.Println(val)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	arr := []int{1, 2, 3, 4, 5, 8}
	in := make(chan int)
	out := make(chan int)
	go mul2(in, out)
	wg.Add(1)
	go read(out, &wg)
	for _, val := range arr {
		in <- val
	}
	close(in)
	wg.Wait()
}
