package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIntersperse(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{}, Intersperse([]int{}, 7))
    assert.Equal([]int{5, 8, 6, 8, 3, 8, 2}, Intersperse([]int{5, 6, 3, 2}, 8))
    assert.Equal([]int{2, 2, 2, 2, 2}, Intersperse([]int{2, 2, 2}, 2))
}
