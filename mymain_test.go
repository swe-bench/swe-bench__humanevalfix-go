package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStringXor(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("010010", StringXor("111000", "101010"))
    assert.Equal("0", StringXor("1", "1"))
    assert.Equal("0101", StringXor("0101", "0000"))
}
