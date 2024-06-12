package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetPositive(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{4, 5, 6}, GetPositive([]int{-1, -2, 4,5, 6}))
    assert.Equal([]int{5, 3, 2, 3, 3, 9, 123, 1}, GetPositive([]int{5, 3, -5, 2, 3, 3, 9, 0, 123, 1, -10}))
    assert.Equal([]int{}, GetPositive([]int{-1, -2}))
    assert.Equal([]int{}, GetPositive([]int{}))
}
