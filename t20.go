package main

import (
	"fmt"
	"strings"
)

// reverseWords flips words in the line
func reverseWords(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}
	left := 0
	right := len(words) - 1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))
}
