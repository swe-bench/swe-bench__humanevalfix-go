package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(55, Fib(10))
    assert.Equal(1, Fib(1))
    assert.Equal(21, Fib(8))
    assert.Equal(89, Fib(11))
    assert.Equal(144, Fib(12))
}
