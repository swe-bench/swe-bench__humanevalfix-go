package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strconv"
)

// Return a string containing space-delimited numbers starting from 0 upto n inclusive.
// >>> StringSequence(0)
// '0'
// >>> StringSequence(5)
// '0 1 2 3 4 5'
func StringSequence(n int) string{

    var seq string
    for i := 0; i < n; i++ {
        seq += strconv.Itoa(i)
        if i != n {
            seq += " "
        }
    }
    return seq
}

func ExampleTestStringSequence(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("0", StringSequence(0))
    assert.Equal("0 1 2 3 4 5", StringSequence(5))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestStringSequence(t)
}