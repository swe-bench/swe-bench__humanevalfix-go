package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// You are given a non-empty list of positive integers. Return the greatest integer that is greater than
// zero, and has a frequency greater than or equal to the value of the integer itself.
// The frequency of an integer is the number of times it appears in the list.
// If no such a value exist, return -1.
// Examples:
// Search([4, 1, 2, 2, 3, 1]) == 2
// Search([1, 2, 2, 3, 3, 3, 4, 4, 4]) == 3
// Search([5, 5, 4, 4, 4]) == -1
func Search(lst []int) int {

    countMap := make(map[int]int)
	for _, i := range lst {
		if count, ok := countMap[i]; ok {
			countMap[i] = count + 1
		} else {
			countMap[i] = 1
		}
	}
	max := 0
	for i, count := range countMap {
		if count >= i && count > max {
			max = i
		}
	}
	return max
}

func ExampleTestSearch(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(2, Search([]int{4, 1, 2, 2, 3, 1}))
    assert.Equal(3, Search([]int{1, 2, 2, 3, 3, 3, 4, 4, 4}))
    assert.Equal(-1, Search([]int{5, 5, 4, 4, 4}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSearch(t)
}