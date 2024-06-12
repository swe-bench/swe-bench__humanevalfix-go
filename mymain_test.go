package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSortArray(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{1, 2, 4, 3, 5}, SortArray([]int{1,5,2,3,4}))
    assert.Equal([]int{-4, -2, -6, -5, -3}, SortArray([]int{-2,-3,-4,-5,-6}))
    assert.Equal([]int{0, 1, 2, 4, 3}, SortArray([]int{1,0,2,3,4}))
    assert.Equal([]int{}, SortArray([]int{}))
    assert.Equal([]int{2, 2, 4, 4, 3, 3, 5, 5, 5, 7, 77}, SortArray([]int{2,5,77,4,5,3,5,7,2,3,4}))
    assert.Equal([]int{32, 3, 5, 6, 12, 44}, SortArray([]int{3,6,44,12,32,5}))
    assert.Equal([]int{2, 4, 8, 16, 32}, SortArray([]int{2,4,8,16,32}))
    assert.Equal([]int{2, 4, 8, 16, 32}, SortArray([]int{2,4,8,16,32}))
}
