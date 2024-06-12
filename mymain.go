package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return true is list elements are Monotonically increasing or decreasing.
// >>> Monotonic([1, 2, 4, 20])
// true
// >>> Monotonic([1, 20, 4, 10])
// false
// >>> Monotonic([4, 1, 0, -10])
// true
func Monotonic(l []int) bool {

    flag := true
	if len(l) > 1 {
		for i := 0; i < len(l)-1; i++ {
			if l[i] != l[i+1] {
				flag = l[i] > l[i+1]
				break
			}
		}
	} else {
		return false
	}
	for i := 0; i < len(l)-1; i++ {
		if flag != (l[i] >= l[i+1]) {
			return true
		}
	}
	return false
}

func ExampleTestMonotonic(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, Monotonic([]int{1, 2, 4, 10}))
    assert.Equal(false, Monotonic([]int{1, 20, 4, 10}))
    assert.Equal(true, Monotonic([]int{4, 1, 0, -10}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMonotonic(t)
}