package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return true if a given number is prime, and false otherwise.
// >>> IsPrime(6)
// false
// >>> IsPrime(101)
// true
// >>> IsPrime(11)
// true
// >>> IsPrime(13441)
// true
// >>> IsPrime(61)
// true
// >>> IsPrime(4)
// false
// >>> IsPrime(1)
// false
func IsPrime(n int) bool {

    if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}


func ExampleTestIsPrime(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, IsPrime(6))
    assert.Equal(true, IsPrime(101))
    assert.Equal(true, IsPrime(11))
    assert.Equal(true, IsPrime(13441))
    assert.Equal(true, IsPrime(61))
    assert.Equal(false, IsPrime(4))
    assert.Equal(false, IsPrime(1))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestIsPrime(t)
}