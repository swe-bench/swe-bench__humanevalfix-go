package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSplitWords(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"Hello", "world!"}, SplitWords("Hello world!"))
    assert.Equal([]string{"Hello", "world!"}, SplitWords("Hello,world!"))
    assert.Equal([]string{"Hello", "world,!"}, SplitWords("Hello world,!"))
    assert.Equal([]string{"Hello,Hello,world", "!"}, SplitWords("Hello,Hello,world !"))
    assert.Equal(3, SplitWords("abcdef"))
    assert.Equal(2, SplitWords("aaabb"))
    assert.Equal(1, SplitWords("aaaBb"))
    assert.Equal(0, SplitWords(""))
}
