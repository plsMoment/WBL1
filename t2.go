package main

import (
	"fmt"
	"sync"
)

func arraySquares(arr []int) []int {
	var wg sync.WaitGroup
	wg.Add(len(arr))
	squares := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		go func(idx int) {
			squares[idx] = arr[idx] * arr[idx]
			wg.Done()
		}(i)
	}
	wg.Wait()
	return squares
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 4}
	fmt.Println(arraySquares(arr))
}
