package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestClosestInteger(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(10, ClosestInteger("10"))
    assert.Equal(15, ClosestInteger("14.5"))
    assert.Equal(-16, ClosestInteger("-15.5"))
    assert.Equal(15, ClosestInteger("15.3"))
    assert.Equal(0, ClosestInteger("0"))
}
