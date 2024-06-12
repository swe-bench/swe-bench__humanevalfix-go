package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return a greatest common divisor of two integers a and b
// >>> GreatestCommonDivisor(3, 5)
// 1
// >>> GreatestCommonDivisor(25, 15)
// 5
func GreatestCommonDivisor(a int,b int) int{

    if b < 2 {
		return a
	}
	var gcd int = 1
	for i := 2; i < b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = a
		}
	}
	return gcd
}

func ExampleTestGreatestCommonDivisor(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, GreatestCommonDivisor(3, 5))
    assert.Equal(5, GreatestCommonDivisor(25, 15))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestGreatestCommonDivisor(t)
}