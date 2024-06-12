package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSortThird(t *testing.T) {
    assert := assert.New(t)
	same := func(src []int, target []int) bool {
		for i := 0; i < len(src); i++ {
			if src[i] != target[i] {
				return false
			}
		}
		return true
	}
	assert.Equal(true, same([]int{1, 2, 3}, SortThird([]int{1, 2, 3})))
	assert.Equal(true, same([]int{1, 3, -5, 2, -3, 3, 5, 0, 123, 9, -10}, SortThird([]int{5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10})))
	assert.Equal(true, same([]int{-10, 8, -12,3, 23, 2, 4, 11, 12, 5}, SortThird([]int{5, 8, -12, 4, 23, 2, 3, 11, 12, -10})))
	assert.Equal(true, same([]int{2, 6, 3, 4, 8, 9, 5}, SortThird([]int{5, 6, 3, 4, 8, 9, 2})))
	assert.Equal(true, same([]int{2, 8, 3, 4, 6, 9, 5}, SortThird([]int{5, 8, 3, 4, 6, 9, 2})))
	assert.Equal(true, same([]int{2, 6, 9, 4, 8, 3, 5}, SortThird([]int{5, 6, 9, 4, 8, 3, 2})))
	assert.Equal(true, same([]int{2, 6, 3, 4, 8, 9, 5, 1}, SortThird([]int{5, 6, 3, 4, 8, 9, 2, 1})))
}
