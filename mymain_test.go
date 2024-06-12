package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDigits(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(5, Digits(5))
    assert.Equal(5, Digits(54))
    assert.Equal(1, Digits(120))
    assert.Equal(5, Digits(5014))
    assert.Equal(315, Digits(98765))
    assert.Equal(2625, Digits(5576543))
    assert.Equal(0, Digits(2468))
}
