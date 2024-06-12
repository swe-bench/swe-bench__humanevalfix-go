package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// Filter an input list of strings only for ones that contain given substring
// >>> FilterBySubstring([], 'a')
// []
// >>> FilterBySubstring(['abc', 'bacd', 'cde', 'array'], 'a')
// ['abc', 'bacd', 'array']
func FilterBySubstring(stringList []string, substring string) []string {

    result := make([]string, 0)
    for _, x := range stringList {
        if strings.Index(substring, x) != -1 {
            result = append(result, x)
        }
    }
    return result
}

func ExampleTestFilterBySubstring(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{}, FilterBySubstring([]string{}, "a"))
    assert.Equal([]string{"abc", "bacd", "array"}, FilterBySubstring([]string{"abc", "bacd", "cde", "array"}, "a"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFilterBySubstring(t)
}