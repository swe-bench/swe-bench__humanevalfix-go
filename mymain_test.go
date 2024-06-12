package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCanArrange(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(3, CanArrange([]int{1,2,4,3,5}))
    assert.Equal(-1, CanArrange([]int{1,2,4,5}))
    assert.Equal(2, CanArrange([]int{1,4,2,5,6,7,8,9,10}))
    assert.Equal(4, CanArrange([]int{4,8,5,7,3}))
    assert.Equal(-1, CanArrange([]int{}))
}
