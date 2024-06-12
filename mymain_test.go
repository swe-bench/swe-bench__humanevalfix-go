package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFindMax(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("string", FindMax([]string{"name", "of", "string"}))
    assert.Equal("enam", FindMax([]string{"name", "enam", "game"}))
    assert.Equal("aaaaaaa", FindMax([]string{"aaaaaaa", "bb", "cc"}))
    assert.Equal("abc", FindMax([]string{"abc", "cba"}))
    assert.Equal("footbott", FindMax([]string{"play", "this", "game", "of", "footbott"}))
    assert.Equal("gonna", FindMax([]string{"we", "are", "gonna", "rock"}))
    assert.Equal("nation", FindMax([]string{"we", "are", "a", "mad", "nation"}))
    assert.Equal("this", FindMax([]string{"this", "is", "a", "prrk"}))
}
