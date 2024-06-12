package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBelowZero(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, BelowZero([]int{}))
    assert.Equal(false, BelowZero([]int{1, 2, -3, 1, 2, -3}))
    assert.Equal(true, BelowZero([]int{1, 2, -4, 5, 6}))
    assert.Equal(false, BelowZero([]int{1, -1, 2, -2, 5, -5, 4, -4}))
    assert.Equal(true, BelowZero([]int{1, -1, 2, -2, 5, -5, 4, -5}))
    assert.Equal(true, BelowZero([]int{1, -2, 2, -2, 5, -5, 4, -4}))
}
