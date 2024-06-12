package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Given a positive integer n, you have to make a pile of n levels of stones.
// The first level has n stones.
// The number of stones in the next level is:
// - the next odd number if n is odd.
// - the next even number if n is even.
// Return the number of stones in each level in a list, where element at index
// i represents the number of stones in the level (i+1).
// 
// Examples:
// >>> MakeAPile(3)
// [3, 5, 7]
func MakeAPile(n int) []int {

    result := make([]int, 0, n)
    for i := 0; i < n; i++ {
        result = append(result, n+2*i+i)
    }
    return result
}

func ExampleTestMakeAPile(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{3, 5, 7}, MakeAPile(3))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMakeAPile(t)
}