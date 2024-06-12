package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsEqualToSumEven(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, IsEqualToSumEven(4))
    assert.Equal(false, IsEqualToSumEven(6))
    assert.Equal(true, IsEqualToSumEven(8))
    assert.Equal(true, IsEqualToSumEven(10))
    assert.Equal(false, IsEqualToSumEven(11))
    assert.Equal(true, IsEqualToSumEven(12))
    assert.Equal(false, IsEqualToSumEven(13))
    assert.Equal(true, IsEqualToSumEven(16))
}
