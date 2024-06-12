package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// For a given number n, find the largest number that divides n evenly, smaller than n
// >>> LargestDivisor(15)
// 5
func LargestDivisor(n int) int {

    for i := n - 1; i > 0; i-- {
		if n - i == 0 {
			return i
		}
	}
	return 0
}

func ExampleTestLargestDivisor(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(5, LargestDivisor(15))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestLargestDivisor(t)
}