package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Write a function that takes a string and returns true if the string
// length is a prime number or false otherwise
// Examples
// PrimeLength('Hello') == true
// PrimeLength('abcdcba') == true
// PrimeLength('kittens') == true
// PrimeLength('orange') == false
func PrimeLength(s string) bool {

    l := len(s)
    if l == 0 || l == 1 {
        return false
    }
    for i := 3; i < l; i++ {
        if l%i == 0 {
            return false
        }
    }
    return true
}

func ExampleTestPrimeLength(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, PrimeLength("Hello"))
    assert.Equal(true, PrimeLength("abcdcba"))
    assert.Equal(true, PrimeLength("kittens"))
    assert.Equal(false, PrimeLength("orange"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestPrimeLength(t)
}