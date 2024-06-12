package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, IsPrime(6))
    assert.Equal(true, IsPrime(101))
    assert.Equal(true, IsPrime(11))
    assert.Equal(true, IsPrime(13441))
    assert.Equal(true, IsPrime(61))
    assert.Equal(false, IsPrime(4))
    assert.Equal(false, IsPrime(1))
    assert.Equal(true, IsPrime(5))
    assert.Equal(true, IsPrime(11))
    assert.Equal(true, IsPrime(17))
    assert.Equal(false, IsPrime(5 * 17))
    assert.Equal(false, IsPrime(11 * 7))
    assert.Equal(false, IsPrime(13441 * 19))
}
