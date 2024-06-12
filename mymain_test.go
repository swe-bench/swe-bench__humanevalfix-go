package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFlipCase(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("", FlipCase(""))
    assert.Equal("hELLO!", FlipCase("Hello!"))
    assert.Equal("tHESE VIOLENT DELIGHTS HAVE VIOLENT ENDS",FlipCase("These violent delights have violent ends"))
}
