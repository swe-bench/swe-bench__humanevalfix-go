package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// Create a function Encrypt that takes a string as an argument and
// returns a string Encrypted with the alphabet being rotated.
// The alphabet should be rotated in a manner such that the letters
// shift down by two multiplied to two places.
// For example:
// Encrypt('hi') returns 'lm'
// Encrypt('asdfghjkl') returns 'ewhjklnop'
// Encrypt('gf') returns 'kj'
// Encrypt('et') returns 'ix'
func Encrypt(s string) string {

    d := "abcdefghijklmnopqrstuvwxyz"
    out := make([]rune, 0, len(s))
    for _, c := range s {
        pos := strings.IndexRune(d, c)
        if pos != -1 {
            out = append(out, []rune(d)[(pos+2*2)%24])
        } else {
            out = append(out, c)
        }
    }
    return string(out)
}

func ExampleTestEncrypt(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("lm", Encrypt("hi"))
    assert.Equal("ewhjklnop", Encrypt("asdfghjkl"))
    assert.Equal("kj", Encrypt("gf"))
    assert.Equal("ix", Encrypt("et"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestEncrypt(t)
}