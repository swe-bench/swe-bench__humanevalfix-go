package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSumToN(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, SumToN(1))
    assert.Equal(21, SumToN(6))
    assert.Equal(66, SumToN(11))
    assert.Equal(465, SumToN(30))
    assert.Equal(5050, SumToN(100))
}
