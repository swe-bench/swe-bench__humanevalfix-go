package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "sort"
)

// Given an array of non-negative integers, return a copy of the given array after sorting,
// you will sort the given array in ascending order if the sum( first index value, last index value) is odd,
// or sort it in descending order if the sum( first index value, last index value) is even.
// 
// Note:
// * don't change the given array.
// 
// Examples:
// * SortArray([]) => []
// * SortArray([5]) => [5]
// * SortArray([2, 4, 3, 0, 1, 5]) => [0, 1, 2, 3, 4, 5]
// * SortArray([2, 4, 3, 0, 1, 5, 6]) => [6, 5, 4, 3, 2, 1, 0]
func SortArray(array []int) []int {

    arr := make([]int, len(array))
    copy(arr, array)
    if len(arr) == 0 {
        return arr
    }
    if (arr[0]+arr[len(arr)-1])%2 != 0 {
        sort.Slice(arr, func(i, j int) bool {
            return arr[i] > arr[j]
        })
    } else {
        sort.Slice(arr, func(i, j int) bool {
            return arr[i] < arr[j]
        })
    }
    return arr
}

func ExampleTestSortArray(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{}, SortArray([]int{}), "Error")
    assert.Equal([]int{5}, SortArray([]int{5}), "Error")
    assert.Equal([]int{0, 1, 2, 3, 4, 5}, SortArray([]int{2, 4, 3, 0, 1, 5}), "Error")
    assert.Equal([]int{6, 5, 4, 3, 2, 1, 0}, SortArray([]int{2, 4, 3, 0, 1, 5, 6}), "Error")
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSortArray(t)
}