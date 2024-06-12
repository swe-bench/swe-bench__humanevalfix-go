package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStringSequence(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("0", StringSequence(0))
    assert.Equal("0 1 2 3", StringSequence(3))
    assert.Equal("0 1 2 3 4 5 6 7 8 9 10", StringSequence(10))
}
