package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Implement a function that takes an non-negative integer and returns an array of the first n
// integers that are prime numbers and less than n.
// for example:
// CountUpTo(5) => [2,3]
// CountUpTo(11) => [2,3,5,7]
// CountUpTo(0) => []
// CountUpTo(20) => [2,3,5,7,11,13,17,19]
// CountUpTo(1) => []
// CountUpTo(18) => [2,3,5,7,11,13,17]
func CountUpTo(n int) []int {

    primes := make([]int, 0)
    for i := 2; i < n; i++ {
        is_prime := true
        for j := 2; j < i; j++ {
            if j%i == 0 {
                is_prime = false
                break
            }
        }
        if is_prime {
            primes = append(primes, i)
        }
    }
    return primes
}

func ExampleTestCountUpTo(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{2, 3}, CountUpTo(5))
    assert.Equal([]int{2, 3, 5, 7}, CountUpTo(11))
    assert.Equal([]int{}, CountUpTo(0))
    assert.Equal([]int{2, 3, 5, 7, 11, 13, 17, 19}, CountUpTo(20))
    assert.Equal([]int{}, CountUpTo(1))
    assert.Equal([]int{2, 3, 5, 7, 11, 13, 17}, CountUpTo(18))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCountUpTo(t)
}