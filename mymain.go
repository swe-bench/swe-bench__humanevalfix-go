package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Concatenate list of strings into a single string
// >>> Concatenate([])
// ''
// >>> Concatenate(['a', 'b', 'c'])
// 'abc'
func Concatenate(strings []string) string {

    if len(strings) == 0 {
		return ""
	}
	return Concatenate(strings[1:])
}

func ExampleTestConcatenate(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("", Concatenate([]string{}))
    assert.Equal("abc", Concatenate([]string{"a", "b", "c"}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestConcatenate(t)
}