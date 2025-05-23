package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, MaxArea([]int{1, 1}))
	assert.Equal(t, 9, MaxArea([]int{1, 2, 3, 1000, 9}))
	assert.Equal(t, 3, MaxArea([]int{3, 6, 1}))
	assert.Equal(t, 7, MaxArea([]int{8, 7, 2, 1}))
	assert.Equal(t, 4, MaxArea([]int{1, 2, 3, 4}))
	assert.Equal(t, 2, MaxArea([]int{1, 2, 1}))
	assert.Equal(t, 200, MaxArea([]int{1, 8, 100, 2, 100, 4, 8, 3, 7}))
	assert.Equal(t, 24, MaxArea([]int{1, 3, 2, 5, 25, 24, 5}))
}

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 6, MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 1, MaxSubArray([]int{1}))
	assert.Equal(t, 23, MaxSubArray([]int{5, 4, -1, 7, 8}))
	assert.Equal(t, 11, MaxSubArray([]int{2, 3, -8, 7, -1, 2, 3}))
	assert.Equal(t, -2, MaxSubArray([]int{-2, -4}))
	assert.Equal(t, 25, MaxSubArray([]int{5, 4, 1, 7, 8}))
	assert.Equal(t, 1, MaxSubArray([]int{-2, 1}))
	assert.Equal(t, -1, MaxSubArray([]int{-2, -1}))
	assert.Equal(t, 1, MaxSubArray([]int{1, -2, -2, -1}))
}
