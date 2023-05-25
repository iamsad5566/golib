package slice_test

import (
	"testing"

	"github.com/iamsad5566/golib/slice"
	"github.com/stretchr/testify/assert"
)

func TestIntSlice(t *testing.T) {
	slInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 1}
	targetInt := 1
	assert.True(t, slice.Contains(slInt, targetInt))
	assert.False(t, slice.Contains(slInt, 0))
}

func TestFloatSlice(t *testing.T) {
	slFloat := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	targetFloat := 1.0
	assert.True(t, slice.Contains(slFloat, targetFloat))
	assert.False(t, slice.Contains(slFloat, 0.0))
}
