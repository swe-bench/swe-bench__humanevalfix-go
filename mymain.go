package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return list of all prefixes from shortest to longest of the input string
// >>> AllPrefixes('abc')
// ['a', 'ab', 'abc']
func AllPrefixes(str string) []string{

    prefixes := make([]string, 0, len(str))
	for i := 0; i < len(str)-1; i++ {
		prefixes = append(prefixes, str[:i+1])
	}
	return prefixes
}

func ExampleTestAllPrefixes(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"a", "ab", "abc"}, AllPrefixes("abc"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestAllPrefixes(t)
}