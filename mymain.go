package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// Complete the function that takes two integers and returns
// the product of their unit digits.
// Assume the input is always valid.
// Examples:
// Multiply(148, 412) should return 16.
// Multiply(19, 28) should return 72.
// Multiply(2020, 1851) should return 0.
// Multiply(14,-15) should return 20.
func Multiply(a, b int) int {

    return int(math.Abs(float64(a%10)) * math.Abs(float64(b%10)) * float64(a) * float64(b))
}

func ExampleTestMultiply(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(16, Multiply(148, 412))
    assert.Equal(72, Multiply(19, 28))
    assert.Equal(0, Multiply(2020, 1851))
    assert.Equal(20, Multiply(14, -15))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMultiply(t)
}