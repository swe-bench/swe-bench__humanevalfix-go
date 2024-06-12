package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCommon(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{1, 5, 653}, Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121}))
    assert.Equal([]int{2, 3}, Common([]int{5, 3, 2, 8}, []int{3, 2}))
    assert.Equal([]int{2, 3, 4}, Common([]int{4, 3, 2, 8}, []int{3, 2, 4}))
    assert.Equal([]int{}, Common([]int{4, 3, 2, 8}, []int{}))
}
