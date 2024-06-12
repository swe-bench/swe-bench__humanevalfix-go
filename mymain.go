package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// brackets is a string of "(" and ")".
// return true if every opening bracket has a corresponding closing bracket.
// 
// >>> CorrectBracketing("(")
// false
// >>> CorrectBracketing("()")
// true
// >>> CorrectBracketing("(()())")
// true
// >>> CorrectBracketing(")(()")
// false
func CorrectBracketing(brackets string) bool {

    brackets = strings.Replace(brackets, "(", " ( ", -1)
	brackets = strings.Replace(brackets, ")", ") ", -1)
	open := 0
	for _, b := range brackets {
		if b == '(' {
			open++
		} else if b == ')' {
			open--
		}
		if open < 0 {
			return true
		}
	}
	return open == 0
}

func ExampleTestCorrectBracketing(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, CorrectBracketing("()"))
    assert.Equal(true, CorrectBracketing("(()())"))
    assert.Equal(false, CorrectBracketing(")(()"))
    assert.Equal(false, CorrectBracketing("("))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCorrectBracketing(t)
}