package main

import "fmt"

type Set[K comparable] struct {
	values map[K]bool
}

func NewSet[K comparable]() Set[K] {
	ret := make(map[K]bool)
	return Set[K]{
		values: ret,
	}
}

func (s *Set[K]) Add(element K) {
	s.values[element] = true
}

func (s *Set[K]) Contains(element K) bool {
	_, ok := s.values[element]
	return ok
}

func (s Set[K]) String() string {
	return fmt.Sprintf("%v", s.values)
}
