package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCheckIfLastCharIsALetter(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, CheckIfLastCharIsALetter("apple"))
    assert.Equal(true, CheckIfLastCharIsALetter("apple pi e"))
    assert.Equal(false, CheckIfLastCharIsALetter("eeeee"))
    assert.Equal(true, CheckIfLastCharIsALetter("A"))
    assert.Equal(false, CheckIfLastCharIsALetter("Pumpkin pie "))
    assert.Equal(false, CheckIfLastCharIsALetter("Pumpkin pie 1"))
    assert.Equal(false, CheckIfLastCharIsALetter(""))
    assert.Equal(false, CheckIfLastCharIsALetter("eeeee e "))
    assert.Equal(false, CheckIfLastCharIsALetter("apple pie"))
    assert.Equal(false, CheckIfLastCharIsALetter("apple pi e "))
}
