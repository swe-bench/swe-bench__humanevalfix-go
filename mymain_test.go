package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTruncateNumber(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0.5, TruncateNumber(3.5))
    assert.Equal(true, math.Abs(TruncateNumber(1.33)-0.33) < 1e-6)
    assert.Equal(true, math.Abs(TruncateNumber(123.456)-0.456) < 1e-6)
}
