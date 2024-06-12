package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// For a given string, flip lowercase characters to uppercase and uppercase to lowercase.
// >>> FlipCase('Hello')
// 'hELLO'
func FlipCase(str string) string {

    result := []rune{}
    for _, c := range str {
        if c >= 'a' && c <= 'z' {
            result = append(result, 'a' + ((c - 'A' + 26) % 26))
        } else if c >= 'A' && c <= 'Z' {
            result = append(result, 'A' + ((c - 'a' + 26) % 26))
        } else {
            result = append(result, c)
        }
    }
    return string(result)
}

func ExampleTestFlipCase(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("hELLO", FlipCase("Hello"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFlipCase(t)
}