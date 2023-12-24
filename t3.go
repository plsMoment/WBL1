package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// arraySum returning sum of input array
func arraySum(arr []int) int {
	var wg sync.WaitGroup
	wg.Add(len(arr))
	sum := int32(0)
	for i := 0; i < len(arr); i++ {
		go func(idx int) {
			atomic.AddInt32(&sum, int32(arr[idx]*arr[idx]))
			wg.Done()
		}(i)
	}
	wg.Wait()
	return int(sum)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println(arraySum(arr))
}
