package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFixSpaces(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("Example", FixSpaces("Example"))
    assert.Equal("Mudasir_Hanif_", FixSpaces("Mudasir Hanif "))
    assert.Equal("Yellow_Yellow__Dirty__Fellow", FixSpaces("Yellow Yellow  Dirty  Fellow"))
    assert.Equal("Exa-mple", FixSpaces("Exa   mple"))
    assert.Equal("-Exa_1_2_2_mple", FixSpaces("   Exa 1 2 2 mple"))
}
