package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestModp(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(3, Modp(3, 5))
    assert.Equal(2, Modp(1101, 101))
    assert.Equal(1, Modp(0, 101))
    assert.Equal(8, Modp(3, 11))
    assert.Equal(1, Modp(100, 101))
    assert.Equal(4, Modp(30, 5))
    assert.Equal(3, Modp(31, 5))
}
