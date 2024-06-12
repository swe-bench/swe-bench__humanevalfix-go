package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// brackets is a string of "<" and ">".
// return true if every opening bracket has a corresponding closing bracket.
// 
// >>> CorrectBracketing("<")
// false
// >>> CorrectBracketing("<>")
// true
// >>> CorrectBracketing("<<><>>")
// true
// >>> CorrectBracketing("><<>")
// false
func CorrectBracketing(brackets string) bool {

    l := len(brackets)
	count := 0
	for index := 0; index < l; index++ {
		if brackets[index] == '>' {
			count++
		} else if brackets[index] == '<' {
			count--
		}
		if count < 0 {
			return false
		}
	}
    if count == 0 {
        return true
    } else {
        return false
    }
}

func ExampleTestCorrectBracketing(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, CorrectBracketing("<>"))
    assert.Equal(true, CorrectBracketing("<<><>>"))
    assert.Equal(false, CorrectBracketing("><<>"))
    assert.Equal(false, CorrectBracketing("<"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCorrectBracketing(t)
}