package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGreatestCommonDivisor(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, GreatestCommonDivisor(3, 7))
    assert.Equal(5, GreatestCommonDivisor(10, 15))
    assert.Equal(7, GreatestCommonDivisor(49, 14))
    assert.Equal(12, GreatestCommonDivisor(144, 60))
}
