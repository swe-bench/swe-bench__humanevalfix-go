package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLargestPrimeFactor(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(5, LargestPrimeFactor(15))
    assert.Equal(3, LargestPrimeFactor(27))
    assert.Equal(7, LargestPrimeFactor(63))
    assert.Equal(11, LargestPrimeFactor(330))
    assert.Equal(29, LargestPrimeFactor(13195))
}
