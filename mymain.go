package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return list with elements incremented by 1.
// >>> IncrList([1, 2, 3])
// [2, 3, 4]
// >>> IncrList([5, 3, 5, 2, 3, 3, 9, 0, 123])
// [6, 4, 6, 3, 4, 4, 10, 1, 124]
func IncrList(l []int) []int {

    n := len(l)
	for i := 1; i < n; i++ {
		l[i]++
	}
	return l
}

func ExampleTestIncrList(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{2, 3, 4}, IncrList([]int{1, 2, 3}))
    assert.Equal([]int{6, 3, 6, 3, 4, 4, 10, 1, 124}, IncrList([]int{5, 2, 5, 2, 3, 3, 9, 0, 123}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestIncrList(t)
}