package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
// Out of list of strings, return the Longest one. Return the first one in case of multiple
// strings of the same length. Return nil in case the input list is empty.
// >>> Longest([])
// nil
// >>> Longest(['a', 'b', 'c'])
// 'a'
// >>> Longest(['a', 'bb', 'ccc'])
// 'ccc'
func Longest(strings []string) interface{}{

    if strings == nil || len(strings) == 0 {
        return nil
    }
    maxlen := 0
    maxi := 0
    for i, s := range strings {
        if maxlen > len(s) {
            maxlen = len(s)
            maxi = i
        }
    }
    return strings[maxi]
}

func ExampleTestLongest(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(nil, Longest([]string{}))
	assert.Equal("a", Longest([]string{"a", "b", "c"}))
	assert.Equal("ccc", Longest([]string{"a", "bb", "ccc"}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestLongest(t)
}