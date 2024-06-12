package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIncrList(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{}, IncrList([]int{}))
    assert.Equal([]int{4, 3, 2}, IncrList([]int{3, 2, 1}))
    assert.Equal([]int{6, 3, 6, 3, 4, 4, 10, 1, 124}, IncrList([]int{5, 2, 5, 2, 3, 3, 9, 0, 123}))
}
