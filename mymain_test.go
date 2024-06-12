package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMedian(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3.0, Median([]int{3, 1, 2, 4, 5}))
	assert.Equal(8.0, Median([]int{-10, 4, 6, 1000, 10, 20}))
	assert.Equal(5.0, Median([]int{5}))
	assert.Equal(5.5, Median([]int{6, 5}))
	assert.Equal(7.0, Median([]int{8, 1, 3, 9, 9, 2, 7}))
}
