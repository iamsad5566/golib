package set

import (
	"errors"
)

type DataType interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Set[T DataType] map[T]T

func (s *Set[T]) Add(num T) {
	(*s)[num] = num
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

func (s *Set[T]) Next() (T, error) {
	for k, num := range *s {
		delete(*s, k)
		return num, nil
	}
	return 0, errors.New("empty")
}
