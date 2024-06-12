package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPluck(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{2, 1}, Pluck([]int{4,2,3}))
    assert.Equal([]int{2, 1}, Pluck([]int{1,2,3}))
    assert.Equal([]int{}, Pluck([]int{}))
    assert.Equal([]int{0, 1}, Pluck([]int{5, 0, 3, 0, 4, 2}))
    assert.Equal([]int{0, 3}, Pluck([]int{1, 2, 3, 0, 5, 3}))
    assert.Equal([]int{4, 1}, Pluck([]int{5, 4, 8, 4 ,8}))
    assert.Equal([]int{6, 1}, Pluck([]int{7, 6, 7, 1}))
    assert.Equal([]int{}, Pluck([]int{7, 9, 7, 1}))
}
