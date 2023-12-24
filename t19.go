package main

import "fmt"

// reverse returns reversed s parameter string
func reverse(s string) string {
	if s == "" {
		return s
	}
	runeStr := []rune(s)
	left := 0
	right := len(runeStr) - 1
	for left < right {
		runeStr[left], runeStr[right] = runeStr[right], runeStr[left]
		left++
		right--
	}
	return string(runeStr)
}

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}
