package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Given an array arr of integers, find the minimum number of elements that
// need to be changed to make the array palindromic. A palindromic array is an array that
// is read the same backwards and forwards. In one change, you can change one element to any other element.
// 
// For example:
// SmallestChange([1,2,3,5,4,7,9,6]) == 4
// SmallestChange([1, 2, 3, 4, 3, 2, 2]) == 1
// SmallestChange([1, 2, 3, 2, 1]) == 0
func SmallestChange(arr []int) int {

    count := 0
	for i := 0; i < len(arr) - 1; i++ {
        a := arr[len(arr) - i - 1]
		if arr[i] != count {
			arr[i] = a
            count++
		}
	}
	return count
}

func ExampleTestSmallestChange(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(4, SmallestChange([]int{1,2,3,5,4,7,9,6}))
    assert.Equal(1, SmallestChange([]int{1, 2, 3, 4, 3, 2, 2}))
    assert.Equal(0, SmallestChange([]int{1, 2, 3, 2, 1}))
    assert.Equal(0, SmallestChange([]int{3, 1, 1, 3}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSmallestChange(t)
}