package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// The FibFib number sequence is a sequence similar to the Fibbonacci sequnece that's defined as follows:
// Fibfib(0) == 0
// Fibfib(1) == 0
// Fibfib(2) == 1
// Fibfib(n) == Fibfib(n-1) + Fibfib(n-2) + Fibfib(n-3).
// Please write a function to efficiently compute the n-th element of the Fibfib number sequence.
// >>> Fibfib(1)
// 0
// >>> Fibfib(5)
// 4
// >>> Fibfib(8)
// 24
func Fibfib(n int) int {

    if n <= 0 {
		return 0
	}
    switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	default:
		return Fibfib(n-1) + Fibfib(n-2) + Fibfib(n-3)
	}
}

func ExampleTestFibfib(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, Fibfib(1))
    assert.Equal(4, Fibfib(5))
    assert.Equal(24, Fibfib(8))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFibfib(t)
}