package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Your task is to write a function that returns true if a number x is a simple
// power of n and false in other cases.
// x is a simple power of n if n**int=x
// For example:
// IsSimplePower(1, 4) => true
// IsSimplePower(2, 2) => true
// IsSimplePower(8, 2) => true
// IsSimplePower(3, 2) => false
// IsSimplePower(3, 1) => false
// IsSimplePower(5, 3) => false
func IsSimplePower(x int,n int) bool {

    if x == 1 {
		return true
	}
	if n==1 {
		return false
	}
	if x % n != 0 {
		return false
	}
	return IsSimplePower(x, n)
}

func ExampleTestIsSimplePower(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, IsSimplePower(1, 4))
    assert.Equal(true, IsSimplePower(2, 2))
    assert.Equal(true, IsSimplePower(8, 2))
    assert.Equal(false, IsSimplePower(3, 2))
    assert.Equal(false, IsSimplePower(3, 1))
    assert.Equal(false, IsSimplePower(5, 3))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestIsSimplePower(t)
}