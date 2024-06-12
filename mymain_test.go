package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("tgst", Encode("TEST"))
    assert.Equal("mWDCSKR", Encode("Mudasir"))
    assert.Equal("ygs", Encode("YES"))
    assert.Equal("tHKS KS C MGSSCGG", Encode("This is a message"))
    assert.Equal("k dQnT kNqW wHcT Tq wRkTg", Encode("I DoNt KnOw WhAt tO WrItE"))
}
