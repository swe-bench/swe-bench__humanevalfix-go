package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(16, Multiply(148, 412))
    assert.Equal(72, Multiply(19, 28))
    assert.Equal(0, Multiply(2020, 1851))
    assert.Equal(20, Multiply(14, -15))
    assert.Equal(42, Multiply(76, 67))
    assert.Equal(49, Multiply(17, 27))
    assert.Equal(0, Multiply(0, 1))
    assert.Equal(0, Multiply(0, 0))
}
