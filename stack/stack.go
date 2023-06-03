package stack

import "github.com/iamsad5566/golib"

type Stack[T golib.DataType] []T

func(s *Stack[T]) Push(d T) {
	*s = append(*s, d)
}

func(s *Stack[T]) Pop() T {
	l := (*s).Len()
	var last T = (*s)[l-1]
	*s = (*s)[:l-1]
	return last
}

func(s *Stack[T]) IsEmpty() bool {
	return (*s).Len() == 0
}

func(s *Stack[T]) Len() int {
	return len(*s)
}

func NewStack[T golib.DataType]() Stack[T] {
	return make(Stack[T], 0)
}