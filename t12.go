package main

import "fmt"

type StrSet struct {
	data map[string]struct{}
}

func (s *StrSet) String() string {
	return fmt.Sprint(s.data)
}

// NewSSet create new set from input array
func NewSSet(arr []string) *StrSet {
	res := &StrSet{make(map[string]struct{})}
	for _, val := range arr {
		res.data[val] = struct{}{}
	}
	return res
}

func main() {
	s := NewSSet([]string{"cat", "cat", "dog", "cat", "tree"})
	fmt.Println(s)
}
