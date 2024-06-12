package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPrimeFib(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(2, PrimeFib(1))
    assert.Equal(3, PrimeFib(2))
    assert.Equal(5, PrimeFib(3))
    assert.Equal(13, PrimeFib(4))
    assert.Equal(89, PrimeFib(5))
    assert.Equal(233, PrimeFib(6))
    assert.Equal(1597, PrimeFib(7))
    assert.Equal(28657, PrimeFib(8))
    assert.Equal(514229, PrimeFib(9))
    assert.Equal(433494437, PrimeFib(10))
}
