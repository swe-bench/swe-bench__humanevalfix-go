package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStrangeSortList(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{1, 4, 2, 3}, StrangeSortList([]int{1,2, 3, 4}))
    assert.Equal([]int{5, 9, 6, 8, 7}, StrangeSortList([]int{5, 6, 7, 8, 9}))
    assert.Equal([]int{1, 5, 2, 4, 3}, StrangeSortList([]int{1, 2, 3, 4, 5}))
    assert.Equal([]int{1, 9, 5, 8, 6, 7}, StrangeSortList([]int{5, 6, 7, 8, 9, 1}))
    assert.Equal([]int{5, 5, 5, 5}, StrangeSortList([]int{5,5, 5, 5}))
    assert.Equal([]int{}, StrangeSortList([]int{}))
    assert.Equal([]int{1, 8, 2, 7, 3, 6, 4, 5}, StrangeSortList([]int{1,2,3,4,5,6,7,8}))
    assert.Equal([]int{-5, 5, -5, 5, 0, 2, 2, 2}, StrangeSortList([]int{0,2,2,2,5,5,-5,-5}))
    assert.Equal([]int{111111}, StrangeSortList([]int{111111}))
}
