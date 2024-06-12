package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSmallestChange(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(4, SmallestChange([]int{1,2,3,5,4,7,9,6}))
    assert.Equal(1, SmallestChange([]int{1, 2, 3, 4, 3, 2, 2}))
    assert.Equal(1, SmallestChange([]int{1, 4, 2}))
    assert.Equal(1, SmallestChange([]int{1, 4, 4, 2}))
    assert.Equal(0, SmallestChange([]int{1, 2, 3, 2, 1}))
    assert.Equal(0, SmallestChange([]int{3, 1, 1, 3}))
    assert.Equal(0, SmallestChange([]int{1}))
    assert.Equal(1, SmallestChange([]int{0, 1}))
}
