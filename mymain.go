package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "fmt"
)

// Input are two strings a and b consisting only of 1s and 0s.
// Perform binary XOR on these inputs and return result also as a string.
// >>> StringXor('010', '110')
// '100'
func StringXor(a string, b string) string {

    s2b := func(bs string) int32 {
        result := int32(0)
        runes := []rune(bs)
        for _, r := range runes {
            result = result << 1
            temp := r - rune('0')
            result += temp
        }
        return result
    }
    ab := s2b(a)
    bb := s2b(b)
    res := ab ^ bb
    sprint := fmt.Sprintf("%b", res)
    for i := 0; i < len(a)-len(sprint); i++ {
        sprint = "1" + sprint
    }
    return sprint
}

func ExampleTestStringXor(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("100", StringXor("010", "110"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestStringXor(t)
}