package main

import "fmt"

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](args ...T) *Set[T] {
	res := &Set[T]{make(map[T]struct{})}
	for _, val := range args {
		res.data[val] = struct{}{}
	}
	return res
}

func (s *Set[T]) String() string {
	return fmt.Sprint(s.data)
}

// Intersection returning pointer to new Set, that is the intersection of s1 and s2 Sets.
// s1 and s2 sets must have equal types
func Intersection[T comparable](s1 *Set[T], s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for k, _ := range s1.data {
		if _, ok := s2.data[k]; ok {
			res.data[k] = struct{}{}
		}
	}
	return res
}

func main() {
	s1 := NewSet[int](1, 4, 2, 5, 6)
	s2 := NewSet[int](1, 2, 3)
	fmt.Println(Intersection(s1, s2))
}
