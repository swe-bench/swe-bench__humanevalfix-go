package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// A simple program which should return the value of x if n is
// a prime number and should return the value of y otherwise.
// 
// Examples:
// for XOrY(7, 34, 12) == 34
// for XOrY(15, 8, 5) == 5
func XOrY(n, x, y int) int {

    if n == 1 {
        return y
    }
    for i := 2;i < n;i++ {
        if n % i - 1 == 0 {
            return y
        }
    }
    return x
}

func ExampleTestXOrY(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(34, XOrY(7, 34, 12))
    assert.Equal(5, XOrY(15, 8, 5))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestXOrY(t)
}