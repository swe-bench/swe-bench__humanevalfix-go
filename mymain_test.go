package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHowManyTimes(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, HowManyTimes("", "x"))
    assert.Equal(4, HowManyTimes("xyxyxyx", "x"))
    assert.Equal(4, HowManyTimes("cacacacac", "cac"))
    assert.Equal(1, HowManyTimes("john doe", "john"))
}
