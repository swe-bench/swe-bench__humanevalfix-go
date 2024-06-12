package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAllPrefixes(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{}, AllPrefixes(""))
    assert.Equal([]string{"a", "as", "asd", "asdf", "asdfg", "asdfgh"}, AllPrefixes("asdfgh"))
    assert.Equal([]string{"W", "WW", "WWW"}, AllPrefixes("WWW"))
}
