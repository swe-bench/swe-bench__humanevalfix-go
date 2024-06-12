package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConcatenate(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("", Concatenate([]string{}))
    assert.Equal("xyz", Concatenate([]string{"x", "y", "z"}))
    assert.Equal("xyzwk", Concatenate([]string{"x", "y","z", "w", "k"}))
}
