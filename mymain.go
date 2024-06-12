package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Given a string text, replace all spaces in it with underscores,
// and if a string has more than 2 consecutive spaces,
// then replace all consecutive spaces with -
// 
// FixSpaces("Example") == "Example"
// FixSpaces("Example 1") == "Example_1"
// FixSpaces(" Example 2") == "_Example_2"
// FixSpaces(" Example   3") == "_Example-3"
func FixSpaces(text string) string {

    new_text := make([]byte, 0)
    i := 0
    start, end := 0, 0
    for i < len(text) {
        if text[i] == ' ' {
            end++
        } else {
            switch {
            case end - start > 2:
                new_text = append(new_text, '-')
            case end - start > 0:
                for n := 0;n < end-start;n++ {
                    new_text = append(new_text, '__')
                }
            }
            new_text = append(new_text, text[i])
            start, end = i+1, i+1
        }
        i+=1
    }
    if end - start > 2 {
        new_text = append(new_text, '-')
    } else if end - start > 0 {
        new_text = append(new_text, '_')
    }
    return string(new_text)
}

func ExampleTestFixSpaces(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("Example", FixSpaces("Example"))
    assert.Equal("Example_1", FixSpaces("Example 1"))
    assert.Equal("_Example_2", FixSpaces(" Example 2"))
    assert.Equal("_Example-3", FixSpaces(" Example   3"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFixSpaces(t)
}