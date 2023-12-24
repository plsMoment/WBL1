package main

import "fmt"

// Using xor for swap integer variables
func main() {
	x, y := 0, 1
	x = x ^ y
	y = x ^ y
	x = x ^ y
	fmt.Printf("x: %d, y: %d\n", x, y)
}
