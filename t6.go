package main

import (
	"context"
	"fmt"
	"time"
)

// byContext is example of stopping goroutine by context
func byContext(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("Goroutine stopped by context")
}

// byChannel is example of stopping goroutine by channel
func byChannel(done <-chan struct{}) {
	<-done
	fmt.Println("Goroutine stopped by channel")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go byContext(ctx)
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	time.Sleep(2 * time.Second)

	done := make(chan struct{})
	go byChannel(done)
	go func() {
		time.Sleep(time.Second)
		done <- struct{}{}
	}()
	time.Sleep(2 * time.Second)
}
