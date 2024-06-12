package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRoundedAvg(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("0b11", RoundedAvg(1, 5))
    assert.Equal("0b1010", RoundedAvg(7, 13))
    assert.Equal("0b1111001011", RoundedAvg(964, 977))
    assert.Equal("0b1111100101", RoundedAvg(996, 997))
    assert.Equal("0b1011000010", RoundedAvg(560, 851))
    assert.Equal("0b101101110", RoundedAvg(185, 546))
    assert.Equal("0b110101101", RoundedAvg(362, 496))
    assert.Equal("0b1001110010", RoundedAvg(350, 902))
    assert.Equal("0b11010111", RoundedAvg(197, 233))
    assert.Equal(-1, RoundedAvg(7, 5))
    assert.Equal(-1, RoundedAvg(5, 1))
    assert.Equal("0b101", RoundedAvg(5, 5))
}
