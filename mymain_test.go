package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsMultiplyPrime(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, IsMultiplyPrime(5))
    assert.Equal(true, IsMultiplyPrime(30))
    assert.Equal(true, IsMultiplyPrime(8))
    assert.Equal(false, IsMultiplyPrime(10))
    assert.Equal(true, IsMultiplyPrime(125))
    assert.Equal(true, IsMultiplyPrime(3 * 5 * 7))
    assert.Equal(false, IsMultiplyPrime(3 * 6 * 7))
    assert.Equal(false, IsMultiplyPrime(9 * 9 * 9))
    assert.Equal(false, IsMultiplyPrime(11 * 9 * 9))
    assert.Equal(true, IsMultiplyPrime(11 * 13 * 7))
}
