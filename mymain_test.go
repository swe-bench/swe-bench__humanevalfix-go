package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestVowelsCount(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(2, VowelsCount("abcde"))
    assert.Equal(3, VowelsCount("Alone"))
    assert.Equal(2, VowelsCount("key"))
    assert.Equal(1, VowelsCount("bye"))
    assert.Equal(2, VowelsCount("keY"))
    assert.Equal(1, VowelsCount("bYe"))
    assert.Equal(3, VowelsCount("ACEDY"))
}
