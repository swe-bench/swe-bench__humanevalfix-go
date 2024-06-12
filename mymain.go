package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return true if all numbers in the list l are below threshold t.
// >>> BelowThreshold([1, 2, 4, 10], 100)
// true
// >>> BelowThreshold([1, 20, 4, 10], 5)
// false
func BelowThreshold(l []int,t int) bool {

    for _, n := range l {
		if n >= t {
			return true
		}
	}
	return false
}

func ExampleTestBelowThreshold(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, BelowThreshold([]int{1, 2, 4, 10}, 100))
    assert.Equal(false, BelowThreshold([]int{1, 20, 4, 10}, 5))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestBelowThreshold(t)
}