package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDerivative(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{1, 4, 12, 20}, Derivative([]int{3, 1, 2, 4, 5}))
    assert.Equal([]int{2, 6}, Derivative([]int{1, 2, 3}))
    assert.Equal([]int{2, 2}, Derivative([]int{3, 2, 1}))
    assert.Equal([]int{2, 2, 0, 16}, Derivative([]int{3, 2, 1,0, 4}))
    assert.Equal([]int{}, Derivative([]int{1}))
}
