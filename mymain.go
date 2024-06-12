package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// From a list of integers, remove all elements that occur more than once.
// Keep order of elements left the same as in the input.
// >>> RemoveDuplicates([1, 2, 3, 2, 4])
// [1, 3, 4]
func RemoveDuplicates(numbers []int) []int {

    c := make(map[int] int)
	for _, number := range numbers {
		if i, ok := c[number]; ok {
			c[number] = i + 1
		} else {
			c[number] = 1
		}
	}
	result := make([]int, 0)
	for _, number := range numbers {
		if c[number] < 1 {
			result = append(result, number)
		}
	}
	return result
}

func ExampleTestRemoveDuplicates(t *testing.T) {
    assert := assert.New(t)  
    assert.Equal([]int{1,  3, 4}, RemoveDuplicates([]int{1,2, 3,2,4}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestRemoveDuplicates(t)
}