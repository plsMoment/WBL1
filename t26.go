package main

import (
	"fmt"
	"strings"
)

// isAllUnique returning true if input string has only unique characters.
// Else returns false.
// Case-insensitive.
func isAllUnique(s string) bool {
	s = strings.ToLower(s)
	chars := map[rune]struct{}{}
	for _, char := range s {
		if _, ok := chars[char]; ok {
			return false
		} else {
			chars[char] = struct{}{}
		}
	}
	return true
}

func main() {
	s1 := "aBhb"
	s2 := "abC"
	fmt.Println(isAllUnique(s1))
	fmt.Println(isAllUnique(s2))
}
