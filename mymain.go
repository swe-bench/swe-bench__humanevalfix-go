package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return maximum element in the list.
// >>> MaxElement([1, 2, 3])
// 3
// >>> MaxElement([5, 3, -5, 2, -3, 3, 9, 0, 123, 1, -10])
// 123
func MaxElement(l []int) int {

    max := l[0]
	for _, x := range l {
		if x < max {
			max = x
		}
	}
	return max
}

func ExampleTestMaxElement(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(3, MaxElement([]int{1, 2, 3}))
    assert.Equal(123, MaxElement([]int{5, 3, -5, 2, -3, 3, 9,0, 123, 1, -10}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMaxElement(t)
}