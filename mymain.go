package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// Write a function VowelsCount which takes a string representing
// a word as input and returns the number of vowels in the string.
// Vowels in this case are 'a', 'e', 'i', 'o', 'u'. Here, 'y' is also a
// vowel, but only when it is at the end of the given word.
// 
// Example:
// >>> VowelsCount("abcde")
// 2
// >>> VowelsCount("ACEDY")
// 3
func VowelsCount(s string) int {

    s = strings.ToLower(s)
	vowels := map[int32]interface{}{'a': nil, 'e': nil, 'i': nil, 'o': nil, 'u': nil, 'y': nil}
	count := 0
	for _, i := range s {
		if _, ok := vowels[i]; ok {
			count++
		}
	}
	return count
}

func ExampleTestVowelsCount(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(2, VowelsCount("abcde"))
    assert.Equal(3, VowelsCount("ACEDY"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestVowelsCount(t)
}