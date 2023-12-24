package main

import (
	"strings"
)

var justString string

func createHugeString(n int) string {
	return strings.Repeat("в", n)
}

// In original code were two main problems.
// First - slice of string may not contain 100 characters if they are not ASCII.
// Second - slice of huge string don't allow GC sweep huge string, therefore we use our memory less efficiently.
func someFunc() {
	v := createHugeString(1 << 10)
	mv := []rune(v)
	justString = string(mv[:100])
}

func main() {
	someFunc()
	println(justString)
}
