package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Given two positive integers a and b, return the even digits between a
// and b, in ascending order.
// 
// For example:
// GenerateIntegers(2, 8) => [2, 4, 6, 8]
// GenerateIntegers(8, 2) => [2, 4, 6, 8]
// GenerateIntegers(10, 14) => []
func GenerateIntegers(a, b int) []int {

    min := func (a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    max := func (a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    lower := max(2, min(a, b))
    upper := min(8, max(a, b))
    ans := make([]int, 0)
    for i := lower;i < upper;i++ {
        if i&1==0 {
            ans = append(ans, i)
        }
    }
    return ans
}

func ExampleTestGenerateIntegers(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{2, 4, 6, 8}, GenerateIntegers(2, 8))
    assert.Equal([]int{2, 4, 6, 8}, GenerateIntegers(8, 2))
    assert.Equal([]int{}, GenerateIntegers(10, 14))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestGenerateIntegers(t)
}