package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMonotonic(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, Monotonic([]int{1, 2, 4, 10}))
    assert.Equal(true, Monotonic([]int{1, 2, 4, 20}))
    assert.Equal(false, Monotonic([]int{1, 20, 4, 10}))
    assert.Equal(true, Monotonic([]int{4, 1, 0, -10}))
    assert.Equal(true, Monotonic([]int{4, 1, 1, 0}))
    assert.Equal(false, Monotonic([]int{1, 2, 3, 2, 5, 60}))
    assert.Equal(true, Monotonic([]int{1, 2, 3, 4, 5, 60}))
    assert.Equal(true, Monotonic([]int{9, 9, 9, 9}))
}
