package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCountUpper(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, CountUpper("aBCdEf"))
    assert.Equal(0, CountUpper("abcdefg"))
    assert.Equal(0, CountUpper("dBBE"))
    assert.Equal(0, CountUpper("B"))
    assert.Equal(1, CountUpper("U"))
    assert.Equal(0, CountUpper(""))
    assert.Equal(2, CountUpper("EEEE"))
}
