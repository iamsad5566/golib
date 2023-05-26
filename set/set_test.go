package set_test

import (
	"testing"

	"github.com/iamsad5566/golib/set"
	"github.com/iamsad5566/golib/slice"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	set := make(set.Set[int])
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, num := range sl {
		set.Add(num)
	}
	assert.True(t, set.Contains(1))

	prevLen := set.Len()
	set.Remove(1)
	afterLen := set.Len()
	assert.True(t, prevLen == afterLen+1)

	for set.Len() > 0 {
		n, _ := set.Next()
		assert.True(t, slice.Contains(sl, n))
	}
}
