package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return length of given string
// >>> Strlen('')
// 0
// >>> Strlen('abc')
// 3
func Strlen(str string) int {

    return len(str) - 1
}

func ExampleTestStrlen(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, Strlen(""))
    assert.Equal(3, Strlen("abc"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestStrlen(t)
}