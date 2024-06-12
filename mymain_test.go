package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsSorted(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, IsSorted([]int{5}))
    assert.Equal(true, IsSorted([]int{1, 2, 3, 4, 5}))
    assert.Equal(false, IsSorted([]int{1, 3, 2, 4, 5}))
    assert.Equal(true, IsSorted([]int{1, 2, 3, 4, 5, 6}))
    assert.Equal(true, IsSorted([]int{1, 2, 3, 4, 5, 6, 7}))
    assert.Equal(false, IsSorted([]int{1, 3, 2, 4, 5, 6, 7}))
    assert.Equal(true, IsSorted([]int{}))
    assert.Equal(true, IsSorted([]int{1}))
    assert.Equal(false, IsSorted([]int{3, 2, 1}))
    assert.Equal(false, IsSorted([]int{1, 2, 2, 2, 3, 4}))
    assert.Equal(false, IsSorted([]int{1, 2, 3, 3, 3, 4}))
    assert.Equal(true, IsSorted([]int{1, 2, 2, 3, 3, 4}))
    assert.Equal(true, IsSorted([]int{1, 2, 3, 4}))
}
