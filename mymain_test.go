package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIntToMiniRoman(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("xix", IntToMiniRoman(19))
    assert.Equal("clii", IntToMiniRoman(152))
    assert.Equal("ccli", IntToMiniRoman(251))
    assert.Equal("cdxxvi", IntToMiniRoman(426))
    assert.Equal("d", IntToMiniRoman(500))
    assert.Equal("i", IntToMiniRoman(1))
    assert.Equal("iv", IntToMiniRoman(4))
    assert.Equal("xliii", IntToMiniRoman(43))
    assert.Equal("xc", IntToMiniRoman(90))
    assert.Equal("xciv", IntToMiniRoman(94))
    assert.Equal("dxxxii", IntToMiniRoman(532))
    assert.Equal("cm", IntToMiniRoman(900))
    assert.Equal("cmxciv", IntToMiniRoman(994))
    assert.Equal("m", IntToMiniRoman(1000))
}
