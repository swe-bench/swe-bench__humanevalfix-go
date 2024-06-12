package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return the largest prime factor of n. Assume n > 1 and is not a prime.
// >>> LargestPrimeFactor(13195)
// 29
// >>> LargestPrimeFactor(2048)
// 2
func LargestPrimeFactor(n int) int {

    isPrime := func(n int) bool {
        for i := 2; i < int(math.Pow(float64(n), 0.5)+1); i++ {
            if n%i == 0 {
                return false
            }
        }
        return true
    }

    largest := 1
    for j := 2; j < n + 1; j++ {
		if n % j == 0 && isPrime(n) {
			if j > largest {
				largest = j
			}
		}
	}
    return largest
}

func ExampleTestLargestPrimeFactor(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(2, LargestPrimeFactor(2048))
    assert.Equal(29, LargestPrimeFactor(13195))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestLargestPrimeFactor(t)
}