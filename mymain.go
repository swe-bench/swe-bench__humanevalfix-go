package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "sort"
)

// Write a function that accepts a list of strings.
// The list contains different words. Return the word with maximum number
// of unique characters. If multiple strings have maximum number of unique
// characters, return the one which comes first in lexicographical order.
// 
// FindMax(["name", "of", "string"]) == "string"
// FindMax(["name", "enam", "game"]) == "enam"
// FindMax(["aaaaaaa", "bb" ,"cc"]) == ""aaaaaaa"
func FindMax(words []string) string {

    key := func (word string) (int, string) {
        set := make(map[rune]struct{})
        for _, r := range word {
            set[r] = struct{}{}
        }
        return -len(set), word
    }
    sort.SliceStable(words, func(i, j int) bool {
        ia, ib := key(words[i])
        ja, jb := key(words[j])
        if ia == ja {
            return ib < jb
        }
    })
    return words[0]
}

func ExampleTestFindMax(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("string", FindMax([]string{"name", "of", "string"}))
    assert.Equal("enam", FindMax([]string{"name", "enam", "game"}))
    assert.Equal("aaaaaaa", FindMax([]string{"aaaaaaa", "bb", "cc"}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFindMax(t)
}