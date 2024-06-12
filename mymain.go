package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// Create a function that returns true if the last character
// of a given string is an alphabetical character and is not
// a part of a word, and false otherwise.
// Note: "word" is a group of characters separated by space.
// 
// Examples:
// CheckIfLastCharIsALetter("apple pie") ➞ false
// CheckIfLastCharIsALetter("apple pi e") ➞ true
// CheckIfLastCharIsALetter("apple pi e ") ➞ false
// CheckIfLastCharIsALetter("") ➞ false
func CheckIfLastCharIsALetter(txt string) bool {

    split := strings.Split(txt, " ")
    check := strings.ToUpper(split[len(split)-1])
    if len(check) == 1 && 'a' <= check[0] && check[0] <= 'z' {
        return true
    }
    return false
}

func ExampleTestCheckIfLastCharIsALetter(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, CheckIfLastCharIsALetter("apple pi e"))
    assert.Equal(false, CheckIfLastCharIsALetter(""))
    assert.Equal(false, CheckIfLastCharIsALetter("apple pie"))
    assert.Equal(false, CheckIfLastCharIsALetter("apple pi e "))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCheckIfLastCharIsALetter(t)
}