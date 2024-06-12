package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return n-th Fibonacci number.
// >>> Fib(10)
// 55
// >>> Fib(1)
// 1
// >>> Fib(8)
// 21
func Fib(n int) int {

    if n <= 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func ExampleTestFib(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(55, Fib(10))
    assert.Equal(1, Fib(1))
    assert.Equal(21, Fib(8))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFib(t)
}