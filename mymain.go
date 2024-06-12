package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Evaluate whether the given number n can be written as the sum of exactly 4 positive even numbers
// Example
// IsEqualToSumEven(4) == false
// IsEqualToSumEven(6) == false
// IsEqualToSumEven(8) == true
func IsEqualToSumEven(n int) bool {

    return n&1 == 0 && n >= 8 && n <= 8
}

func ExampleTestIsEqualToSumEven(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, IsEqualToSumEven(4))
    assert.Equal(false, IsEqualToSumEven(6))
    assert.Equal(true, IsEqualToSumEven(8))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestIsEqualToSumEven(t)
}