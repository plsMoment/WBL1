package main

import (
	"cmp"
	"fmt"
	"math/rand"
)

// quicksort returns sorted input array.
// Sorts in non-descending order.
// Input array type must be ordered
func quicksort[T cmp.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	low, high := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[high] = arr[high], arr[pivot]

	for i := range arr {
		if arr[i] < arr[high] {
			arr[low], arr[i] = arr[i], arr[low]
			low++
		}
	}

	arr[low], arr[high] = arr[high], arr[low]

	quicksort(arr[:low])
	quicksort(arr[low+1:])

	return arr
}

func main() {
	arr := []int{1, 4, 4, 2, 7, 2, 3}
	fmt.Println(quicksort(arr))
}
