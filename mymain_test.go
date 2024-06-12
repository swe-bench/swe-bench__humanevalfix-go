package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCycpatternCheck(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, CycpatternCheck("xyzw", "xyw"))
    assert.Equal(true, CycpatternCheck("yello", "ell"))
    assert.Equal(false, CycpatternCheck("whattup", "ptut"))
    assert.Equal(true, CycpatternCheck("efef", "fee"))
    assert.Equal(false, CycpatternCheck("abab", "aabb"))
    assert.Equal(true, CycpatternCheck("winemtt", "tinem"))
}
