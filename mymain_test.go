package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSpecialFilter(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, Specialfilter([]int{5, -2, 1, -5}))
    assert.Equal(1, Specialfilter([]int{15, -73, 14, -15}))
    assert.Equal(2, Specialfilter([]int{33, -2, -3, 45, 21, 109}))
    assert.Equal(4, Specialfilter([]int{43, -12, 93, 125, 121, 109}))
    assert.Equal(3, Specialfilter([]int{71, -2, -33, 75, 21, 19}))
    assert.Equal(0, Specialfilter([]int{1}))
    assert.Equal(0, Specialfilter([]int{}))
}               
