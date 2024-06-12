package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// You are given a string s.
// if s[i] is a letter, reverse its case from lower to upper or vise versa,
// otherwise keep it as it is.
// If the string contains no letters, reverse the string.
// The function should return the resulted string.
// Examples
// Solve("1234") = "4321"
// Solve("ab") = "AB"
// Solve("#a@C") = "#A@c"
func Solve(s string) string {

    flg := 0
    new_str := []rune(s)
    for i, r := range new_str {
        if ('a' <= r && r <= 'z') {
            if 'a' <= r && r <= 'z' {
                new_str[i] = r - 'a' + 'A'
            } else {
                new_str[i] = r - 'A' + 'a'
            }
            flg = 1
        }
    }
    if flg == 0 {
        for i := 0;i < len(new_str)>>1;i++ {
            new_str[i], new_str[len(new_str)-i-1] = new_str[len(new_str)-i-1], new_str[i]
        }
    }
    return string(new_str)
}

func ExampleTestSolve(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("4321", Solve("1234"))
    assert.Equal("AB", Solve("ab"))
    assert.Equal("#A@c", Solve("#a@C"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSolve(t)
}