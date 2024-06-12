package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// Given a string, find out how many distinct characters (regardless of case) does it consist of
// >>> CountDistinctCharacters('xyzXYZ')
// 3
// >>> CountDistinctCharacters('Jerry')
// 4
func CountDistinctCharacters(str string) int{

    count := 0
	set := make(map[rune]bool)
	for _, i := range str {
		if set[i] == true {
			continue
		} else {
			set[i] = true
			count++
		}
	}
	return count
}


func ExampleTestCountDistinctCharacters(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(3, CountDistinctCharacters("xyzXYZ"))
    assert.Equal(4, CountDistinctCharacters("Jerry"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCountDistinctCharacters(t)
}