package main

import (
	"fmt"
	"time"
)

// N seconds to finish program
const N = 5

func main() {
	ch := make(chan string)

	go func() {
		for data := range ch {
			fmt.Println(data)
		}
		fmt.Println("Channel closed")
	}()

	timer := time.After(N * time.Second)

	for {
		select {
		case <-timer:
			close(ch)
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Program finished")
			return
		default:
			ch <- "Any data"
		}
		time.Sleep(100 * time.Millisecond)
	}
}
