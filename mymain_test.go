package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("lm", Encrypt("hi"))
    assert.Equal("ewhjklnop", Encrypt("asdfghjkl"))
    assert.Equal("kj", Encrypt("gf"))
    assert.Equal("ix", Encrypt("et"))
    assert.Equal("jeiajeaijeiak", Encrypt("faewfawefaewg"))
    assert.Equal("lippsqcjvmirh", Encrypt("hellomyfriend"))
    assert.Equal("hbdhpqrmpjylqmpyjlpmlyjrqpmqryjlpmqryjljygyjl", Encrypt("dxzdlmnilfuhmilufhlihufnmlimnufhlimnufhfucufh"))
    assert.Equal("e", Encrypt("a"))
}
