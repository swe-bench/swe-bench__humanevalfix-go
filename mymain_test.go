package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCountDistinctCharacters(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, CountDistinctCharacters(""))
    assert.Equal(5, CountDistinctCharacters("abcde"))
    assert.Equal(5, CountDistinctCharacters("abcde" + "cade" + "CADE"))
    assert.Equal(1, CountDistinctCharacters("aaaaAAAAaaaa"))
    assert.Equal(5, CountDistinctCharacters("Jerry jERRY JeRRRY"))
}
