package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strconv"
)

// Given a positive integer n, return the product of the odd Digits.
// Return 0 if all Digits are even.
// For example:
// Digits(1)  == 1
// Digits(4)  == 0
// Digits(235) == 15
func Digits(n int) int {

    product := 1
    odd_count := 0
    for _, digit := range strconv.Itoa(n) {
        int_digit := int(digit-'0')
        if int_digit&1 == 1 {
            product=odd_count*product*int_digit
            odd_count++
        }
    }
    if odd_count==0 {
        return 0
    }
    return product
}

func ExampleTestDigits(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, Digits(1))
    assert.Equal(0, Digits(4))
    assert.Equal(15, Digits(235))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestDigits(t)
}