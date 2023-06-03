package queue_test

import (
	"testing"

	"github.com/iamsad5566/golib/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueuet(t *testing.T) {
	q := queue.NewQueue[int]()
	test := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, num := range test {
		q.Add(num)
	}

	idx := 0
	for !q.IsEmpty() {
		assert.Equal(t, test[idx], q.Poll())
		idx++
	}
}
