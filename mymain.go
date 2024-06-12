package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// You are given 2 words. You need to return true if the second word or any of its rotations is a substring in the first word
// CycpatternCheck("abcd","abd") => false
// CycpatternCheck("hello","ell") => true
// CycpatternCheck("whassup","psus") => false
// CycpatternCheck("abab","baa") => true
// CycpatternCheck("efef","eeff") => false
// CycpatternCheck("himenss","simen") => true
func CycpatternCheck(a , b string) bool {

    l := len(b)
    pat := b + b
    for i := 0;i < len(a) - l + 1; i++ {
        for j := 0;j < len(b) - l + 1;j++ {
            if a[i:i+l] == pat[j:j+l] {
                return true
            }
        }
    }
    return false
}

func ExampleTestCycpatternCheck(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, CycpatternCheck("abcd", "abd"))
    assert.Equal(true, CycpatternCheck("hello", "ell"))
    assert.Equal(false, CycpatternCheck("whasup", "psus"))
    assert.Equal(true, CycpatternCheck("abab", "baa"))
    assert.Equal(false, CycpatternCheck("efef", "eeff"))
    assert.Equal(true, CycpatternCheck("himenss", "simen"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCycpatternCheck(t)
}