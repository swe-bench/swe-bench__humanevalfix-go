package main

import (
    "testing"
    "math"
    "github.com/stretchr/testify/assert"
)

func TestCountNums(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, CountNums([]int{}))
    assert.Equal(0, CountNums([]int{-1, -2, 0}))
    assert.Equal(6, CountNums([]int{1, 1, 2, -2, 3, 4, 5}))
    assert.Equal(5, CountNums([]int{1, 6, 9, -6, 0, 1, 5}))
    assert.Equal(4, CountNums([]int{1, 100, 98, -7, 1, -1}))
    assert.Equal(5, CountNums([]int{12, 23, 34, -45, -56, 0}))
    assert.Equal(1, CountNums([]int{-0, int(math.Pow(1, 0))}))
    assert.Equal(1, CountNums([]int{1}))
}
