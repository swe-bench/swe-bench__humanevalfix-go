package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// Given a positive floating point number, it can be decomposed into
// and integer part (largest integer smaller than given number) and decimals
// (leftover part always smaller than 1).
// 
// Return the decimal part of the number.
// >>> TruncateNumber(3.5)
// 0.5
func TruncateNumber(number float64) float64 {

    return math.Mod(number,1)+1
}

func ExampleTestTruncateNumber(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0.5, TruncateNumber(3.5))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestTruncateNumber(t)
}