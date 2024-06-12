package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMaximum(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{-4, -3, 5}, Maximum([]int{-3, -4, 5}, 3))
    assert.Equal([]int{4, 4}, Maximum([]int{4, -4, 4}, 2))
    assert.Equal([]int{2}, Maximum([]int{-3, 2, 1, 2, -1, -2, 1}, 1))
    assert.Equal([]int{2, 20, 123}, Maximum([]int{123, -123, 20, 0 , 1, 2, -3}, 3))
    assert.Equal([]int{0, 1, 2, 20}, Maximum([]int{-123, 20, 0 , 1, 2, -3}, 4))
    assert.Equal([]int{-13, -8, 0, 0, 3, 5, 15}, Maximum([]int{5, 15, 0, 3, -13, -8, 0}, 7))
    assert.Equal([]int{3, 5}, Maximum([]int{-1, 0, 2, 5, 3, -10}, 2))
    assert.Equal([]int{5}, Maximum([]int{1, 0, 5, -7}, 1))
    assert.Equal([]int{-4, 4}, Maximum([]int{4, -4}, 2))
    assert.Equal([]int{-10, 10}, Maximum([]int{-10, 10}, 2))
    assert.Equal([]int{}, Maximum([]int{1, 2, 3, -23, 243, -400, 0}, 0))
  }
