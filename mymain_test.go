package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFilterBySubstring(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{}, FilterBySubstring([]string{}, "john"))
    assert.Equal([]string{"xxx", "xxxAAA", "xxx"}, FilterBySubstring([]string{"xxx", "asd", "xxy", "john doe", "xxxAAA", "xxx"}, "xxx"))
    assert.Equal([]string{"xxx", "aaaxxy", "xxxAAA", "xxx"}, FilterBySubstring([]string{"xxx", "asd", "aaaxxy", "john doe", "xxxAAA", "xxx"}, "xx"))
    assert.Equal([]string{"grunt", "prune"}, FilterBySubstring([]string{"grunt", "trumpet", "prune", "gruesome"}, "run"))
}
