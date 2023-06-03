package set

import (
	"github.com/iamsad5566/golib"
)

type Set[T golib.DataType] map[T]T

func (s *Set[T]) Add(num T) {
	(*s)[num] = num
}

func (s *Set[T]) Contains(num T) bool {
	_, isExisted := (*s)[num]
	return isExisted
}

func (s *Set[T]) Remove(num T) bool {
	if _, ok := (*s)[num]; ok {
		delete(*s, num)
		return true
	}
	return false
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) Next() *T {
	for k, num := range *s {
		delete(*s, k)
		return &num
	}
	return nil
}

func NewSet[T golib.DataType]() Set[T] {
	set := make(Set[T], 0)
	return set
}
