package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSumProduct(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([2]int{0, 1}, SumProduct([]int{}))
    assert.Equal([2]int{3, 1}, SumProduct([]int{1, 1, 1}))
    assert.Equal([2]int{100, 0}, SumProduct([]int{100, 0}))
    assert.Equal([2]int{3 + 5 + 7, 3 * 5 * 7}, SumProduct([]int{3, 5, 7}))
    assert.Equal([2]int{10, 10}, SumProduct([]int{10}))
}
