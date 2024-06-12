package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// SumToN is a function that sums numbers from 1 to n.
// >>> SumToN(30)
// 465
// >>> SumToN(100)
// 5050
// >>> SumToN(5)
// 15
// >>> SumToN(10)
// 55
// >>> SumToN(1)
// 1
func SumToN(n int) int {

    if n <= 0 {
		return 0
	} else {
		return 1 + SumToN(n - 1)
	}
}

func ExampleTestSumToN(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, SumToN(1))
    assert.Equal(15, SumToN(5))
    assert.Equal(55, SumToN(10))
    assert.Equal(465, SumToN(30))
    assert.Equal(5050, SumToN(100))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSumToN(t)
}