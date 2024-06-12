package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLongest(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(nil, Longest([]string{}))
	assert.Equal("x", Longest([]string{"x", "y", "z"}))
	assert.Equal("zzzz", Longest([]string{"x", "yyy", "zzzz", "www", "kkkk", "abc"}))
}
