package main

import (
    "testing"
    "math"
    "github.com/stretchr/testify/assert"
)

func TestNextSmallest(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(2, NextSmallest([]int{1, 2, 3, 4, 5}))
    assert.Equal(2, NextSmallest([]int{5, 1, 4, 3, 2}))
    assert.Equal(nil, NextSmallest([]int{}))
    assert.Equal(nil, NextSmallest([]int{1, 1}))
    assert.Equal(1, NextSmallest([]int{1,1,1,1,0}))
    assert.Equal(nil, NextSmallest([]int{1, int(math.Pow(0, 0))}))
    assert.Equal(-35, NextSmallest([]int{-35, 34, 12, -45}))
}
