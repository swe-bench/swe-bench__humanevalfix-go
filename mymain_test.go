package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("YES", Exchange([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))
    assert.Equal("NO", Exchange([]int{1, 2, 3, 4}, []int{1, 5, 3, 4}))
    assert.Equal("YES", Exchange([]int{1, 2, 3, 4}, []int{2, 1, 4, 3}))
    assert.Equal("YES", Exchange([]int{5, 7, 3}, []int{2, 6, 4}))
    assert.Equal("NO", Exchange([]int{5, 7, 3}, []int{2, 6, 3}))
    assert.Equal("NO", Exchange([]int{3, 2, 6, 1, 8, 9}, []int{3, 5, 5, 1, 1, 1}))
    assert.Equal("YES", Exchange([]int{100, 200}, []int{200, 200}))
}
