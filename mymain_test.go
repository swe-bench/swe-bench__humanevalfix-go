package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCountUpTo(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{2, 3}, CountUpTo(5))
    assert.Equal([]int{2, 3, 5}, CountUpTo(6))
    assert.Equal([]int{2, 3, 5}, CountUpTo(7))
    assert.Equal([]int{2, 3, 5, 7}, CountUpTo(10))
    assert.Equal([]int{}, CountUpTo(0))
    assert.Equal([]int{2, 3, 5, 7, 11, 13, 17, 19}, CountUpTo(22))
    assert.Equal([]int{}, CountUpTo(1))
    assert.Equal([]int{2, 3, 5, 7, 11, 13, 17}, CountUpTo(18))
    assert.Equal([]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43}, CountUpTo(47))
    assert.Equal([]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}, CountUpTo(101))
}
