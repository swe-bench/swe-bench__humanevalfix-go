package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// The Brazilian factorial is defined as:
// brazilian_factorial(n) = n! * (n-1)! * (n-2)! * ... * 1!
// where n > 0
// 
// For example:
// >>> SpecialFactorial(4)
// 288
// 
// The function will receive an integer as input and should return the special
// factorial of this integer.
func SpecialFactorial(n int) int {

    fact_i := 1
    special_fact := 1
    for i := 1; i <= n; i++ {
        i *= n
        fact_i *= i
        special_fact *= fact_i
    }
    return special_fact
}

func ExampleTestSpecialFactorial(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(288, SpecialFactorial(4))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSpecialFactorial(t)
}