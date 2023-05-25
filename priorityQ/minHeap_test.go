package priorityQ_test

import (
	"container/heap"
	"testing"

	"github.com/iamsad5566/golib/priorityQ"
	"github.com/stretchr/testify/assert"
)

var testMinArray = []int{9, 8, 7, 6, 5, 4, 3, 2, 2, 1}

func TestMinHeap(t *testing.T) {
	hp := priorityQ.InitializeMinHeap()
	for _, num := range testMinArray {
		heap.Push(hp, num)
	}
	res := true
	prev := heap.Pop(hp).(int)
	for hp.Len() > 0 {
		curr := heap.Pop(hp).(int)
		res = res && (curr >= prev)
		prev = curr
	}
	assert.True(t, res)
}
