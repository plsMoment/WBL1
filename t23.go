package main

import (
	"fmt"
	"slices"
)

// deleteElem removes the element idx from arr, returning the modified slice
func deleteElem[S ~[]T, T any](arr S, idx uint) S {
	return append(arr[:idx], arr[idx+1:]...)
}

// Implemented in Go same as above function deleteElem
func builtinDeleteElem[S ~[]T, T any](arr S, idx int) S {
	return slices.Delete(arr, idx, idx+1)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(deleteElem(arr, 5))
	fmt.Println(builtinDeleteElem(arr, 5))
}
