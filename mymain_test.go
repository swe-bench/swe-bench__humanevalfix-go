package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRemoveVowels(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("", RemoveVowels(""))
    assert.Equal("bcdf\nghjklm", RemoveVowels("abcdef\nghijklm"))
    assert.Equal("fdcb", RemoveVowels("fedcba"))
    assert.Equal("", RemoveVowels("eeeee"))
    assert.Equal("cB", RemoveVowels("acBAA"))
    assert.Equal("cB", RemoveVowels("EcBOO"))
    assert.Equal("ybcd", RemoveVowels("ybcd"))
}
