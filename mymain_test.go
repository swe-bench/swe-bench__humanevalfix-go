package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSortEven(t *testing.T) {
	assert := assert.New(t)
	same := func(src []int, target []int) bool {
		for i := 0; i < len(src); i++ {
			if src[i] != target[i] {
				return false
			}
		}
		return true
	}
	assert.Equal(true, same([]int{1, 2, 3}, SortEven([]int{1, 2, 3})))
	assert.Equal(true, same([]int{-10, 3, -5, 2, -3, 3, 5, 0, 9, 1, 123}, SortEven([]int{5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10})))
	assert.Equal(true, same([]int{-12, 8, 3, 4, 5, 2, 12, 11, 23, -10}, SortEven([]int{5, 8, -12, 4, 23, 2, 3, 11, 12, -10})))
}
