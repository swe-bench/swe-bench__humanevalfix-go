package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, FizzBuzz(50))
    assert.Equal(2, FizzBuzz(78))
    assert.Equal(3, FizzBuzz(79))
    assert.Equal(3, FizzBuzz(100))
    assert.Equal(6, FizzBuzz(200))
    assert.Equal(192, FizzBuzz(4000))
    assert.Equal(639, FizzBuzz(10000))
    assert.Equal(8026, FizzBuzz(100000))
}
