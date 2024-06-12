package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFibfib(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, Fibfib(2))
    assert.Equal(0, Fibfib(1))
    assert.Equal(4, Fibfib(5))
    assert.Equal(24, Fibfib(8))
    assert.Equal(81, Fibfib(10))
    assert.Equal(274, Fibfib(12))
    assert.Equal(927, Fibfib(14))
}
