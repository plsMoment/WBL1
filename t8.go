package main

import "fmt"

// setBitTo устанавливает idx бит числа val в число setTo
func setBitTo(val int64, idx uint8, setTo uint8) int64 {
	if setTo == 0 {
		val &^= 1 << idx
	} else if setTo == 1 {
		val |= 1 << idx
	} else {
		panic("setTo value must be 0 or 1")
	}
	return val
}

func main() {
	var val int64 = 0
	fmt.Println(setBitTo(val, 63, 1))
}
