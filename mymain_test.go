package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLargestDivisor(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, LargestDivisor(3))
    assert.Equal(1, LargestDivisor(7))
    assert.Equal(5, LargestDivisor(10))
    assert.Equal(50, LargestDivisor(100))
    assert.Equal(7, LargestDivisor(49))
}
