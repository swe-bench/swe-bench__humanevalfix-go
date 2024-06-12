package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAddElements(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(-4, AddElements([]int{1,-2,-3,41,57,76,87,88,99}, 3))
    assert.Equal(0, AddElements([]int{111,121,3,4000,5,6}, 2))
    assert.Equal(125, AddElements([]int{11,21,3,90,5,6,7,8,9}, 4))
    assert.Equal(24, AddElements([]int{111,21,3,4000,5,6,7,8,9}, 4))
    assert.Equal(1, AddElements([]int{1}, 1))
}
