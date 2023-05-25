package priorityQ_test

import (
	"container/heap"
	"testing"

	"github.com/iamsad5566/golib/priorityQ"
	"github.com/stretchr/testify/assert"
)

var testMaxArray = []int{1, 2, 3, 4, 5, 6, 7, 8}

func TestMaxHeap(t *testing.T) {
	hp := priorityQ.InitializeMaxHeap()
	for _, num := range testMaxArray {
		heap.Push(hp, num)
	}
	res := true
	prev := heap.Pop(hp).(int)
	for hp.Len() > 0 {
		curr := heap.Pop(hp).(int)
		res = res && (curr <= prev)
		prev = curr
	}
	assert.True(t, res)
}
