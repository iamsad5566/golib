package stack_test

import (
	"testing"

	"github.com/iamsad5566/golib/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := stack.NewStack[int]()
	test := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, num := range test {
		stack.Push(num)
	}

	idx := len(test)-1
	for !stack.IsEmpty() {
		assert.Equal(t, stack.Pop(), test[idx])
		idx--
	}
}