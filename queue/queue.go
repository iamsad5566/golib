package queue

import "github.com/iamsad5566/golib"

type Queue[T golib.DataType] []T

func (q *Queue[T]) Add(num T) {
	*q = append(*q, num)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Poll() T {
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func NewQueue[T golib.DataType]() Queue[T] {
	q := make(Queue[T], 0)
	return q
}
