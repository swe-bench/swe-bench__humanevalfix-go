package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{0,0,0,0,3,3}, Compare([]int{1,2,3,4,5,1}, []int{1,2,3,4,2,-2}))
    assert.Equal([]int{4,4,1,0,0,6}, Compare([]int{0,5,0,0,0,4}, []int{4,1,1,0,0,-2}))
    assert.Equal([]int{0,0,0,0,3,3}, Compare([]int{1,2,3,4,5,1}, []int{1,2,3,4,2,-2}))
    assert.Equal([]int{0,0,0,0,0,0}, Compare([]int{0,0,0,0,0,0}, []int{0,0,0,0,0,0}))
    assert.Equal([]int{2,4,6}, Compare([]int{1,2,3}, []int{-1,-2,-3}))
    assert.Equal([]int{2,0,0,1}, Compare([]int{1,2,3,5}, []int{-1,2,3,4}))
}
