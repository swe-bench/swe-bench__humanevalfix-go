package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMaxElement(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(3, MaxElement([]int{1, 2, 3}))
    assert.Equal(124, MaxElement([]int{5, 3, -5, 2, -3, 3, 9,0, 124, 1, -10}))
}
