package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// workersJob function implements WorkerPool pattern
func workersJob(ctx context.Context, numberOfWorkers int, ch <-chan int, wg *sync.WaitGroup) {
	wg.Add(numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		go func(worker int) {
			defer wg.Done()
			for {
				select {
				case data := <-ch:
					fmt.Printf("Worker %d gets data: %d\n", worker, data)
				case <-ctx.Done():
					fmt.Printf("Worker %d finished\n", worker)
					return
				}
			}
		}(i)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	numberOfWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("error during parse first argument")
	}
	if numberOfWorkers < 1 {
		log.Fatal("first argument(number of workers) must be more than 0")
	}

	var wg sync.WaitGroup
	bufCh := make(chan int, numberOfWorkers)
	counter := 0
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	workersJob(ctx, numberOfWorkers, bufCh, &wg)

	for {
		select {
		case <-done:
			cancel()
			wg.Wait()
			close(bufCh)
			fmt.Println("Program interrupted")
			return
		default:
			bufCh <- counter
		}
		counter++
		time.Sleep(500 * time.Millisecond)
	}
}
