package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMakeAPile(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{3, 5, 7}, MakeAPile(3))
    assert.Equal([]int{4, 6, 8, 10}, MakeAPile(4))
    assert.Equal([]int{5, 7, 9, 11, 13}, MakeAPile(5))
    assert.Equal([]int{6, 8, 10, 12, 14, 16}, MakeAPile(6))
    assert.Equal([]int{8, 10, 12, 14, 16, 18, 20, 22}, MakeAPile(8))
}
