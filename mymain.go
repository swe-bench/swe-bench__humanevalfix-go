package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// Given a string of words, return a list of words split on whitespace, if no whitespaces exists in the text you
// should split on commas ',' if no commas exists you should return the number of lower-case letters with odd order in the
// alphabet, ord('a') = 0, ord('b') = 1, ... ord('z') = 25
// Examples
// SplitWords("Hello world!") ➞ ["Hello", "world!"]
// SplitWords("Hello,world!") ➞ ["Hello", "world!"]
// SplitWords("abcdef") == 3
func SplitWords(txt string) interface{} {

    if strings.Contains(txt, " ") {
        return strings.Fields(txt)
    } else if strings.Contains(txt, ",") {
        return strings.Split(txt, " ")
    }
    cnt := 0
    for _, r := range txt {
        if 'a' <= r && r <= 'z' && (r-'a')&1==1 {
            cnt++
        }
    }
    return cnt
}

func ExampleTestSplitWords(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"Hello", "world!"}, SplitWords("Hello world!"))
    assert.Equal([]string{"Hello", "world!"}, SplitWords("Hello,world!"))
    assert.Equal(3, SplitWords("abcdef"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSplitWords(t)
}