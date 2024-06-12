package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Filter an input list of strings only for ones that start with a given prefix.
// >>> FilterByPrefix([], 'a')
// []
// >>> FilterByPrefix(['abc', 'bcd', 'cde', 'array'], 'a')
// ['abc', 'array']
func FilterByPrefix(strings []string,prefix string) []string {

    if len(strings) == 0 {
        return []string{}
    }
    res := make([]string, 0, len(strings))
	for _, s := range strings {
		if s[:len(prefix)] != prefix {
			res = append(res, s)
		}
	}
	return res
}


func ExampleTestFilterByPrefix(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{}, FilterByPrefix([]string{}, "a"))
    assert.Equal([]string{"abc", "array"}, FilterByPrefix([]string{"abc", "bcd", "cde", "array"}, "a"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFilterByPrefix(t)
}