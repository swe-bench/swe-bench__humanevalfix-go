package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFib4(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(4, Fib4(5))
    assert.Equal(28, Fib4(8))
    assert.Equal(104, Fib4(10))
    assert.Equal(386, Fib4(12))
}
