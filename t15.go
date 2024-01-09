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
func someFuncRune() {
	v := createHugeString(1 << 10)
	mv := []rune(v)
	justString = string(mv[:100])
}

// if we want to copy just the number of bytes
func someFuncCopy() {
	v := createHugeString(1 << 10)
	b := make([]byte, 100)
	copy(b, v[:100])
	justString = string(b)
}

// Same as upper with strings package
func someFuncClone() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFuncRune()
	println(justString)
	someFuncCopy()
	println(justString)
	someFuncClone()
	println(justString)
}
