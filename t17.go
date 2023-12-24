package main

import (
	"cmp"
	"fmt"
)

// binSearch returns index of val value in arr array.
// If such element doesn't exist returns -1.
// Argument arr must be in non-descending order and permits any ordered type
func binSearch[T cmp.Ordered](arr []T, val T) int {
	if len(arr) == 0 {
		panic("empty array")
	}
	left := 0
	right := len(arr) - 1
	var mid int
	for left <= right {
		mid = (right-left)/2 + left
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arrNums := []int{1, 1, 2, 3, 4}
	arrStrs := []string{"a", "b", "b", "c"}
	fmt.Println(binSearch(arrNums, 1))
	fmt.Println(binSearch(arrStrs, "b"))
}
