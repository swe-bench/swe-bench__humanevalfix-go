package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "sort"
)

// Given an array arr of integers and a positive integer k, return a sorted list
// of length k with the Maximum k numbers in arr.
// 
// Example 1:
// 
// Input: arr = [-3, -4, 5], k = 3
// Output: [-4, -3, 5]
// 
// Example 2:
// 
// Input: arr = [4, -4, 4], k = 2
// Output: [4, 4]
// 
// Example 3:
// 
// Input: arr = [-3, 2, 1, 2, -1, -2, 1], k = 1
// Output: [2]
// 
// Note:
// 1. The length of the array will be in the range of [1, 1000].
// 2. The elements in the array will be in the range of [-1000, 1000].
// 3. 0 <= k <= len(arr)
func Maximum(arr []int, k int) []int {

    if k == 0 {
        return []int{}
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j] + k
    })
    return arr[len(arr)-k:]
}

func ExampleTestMaximum(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{-4, -3, 5}, Maximum([]int{-3, -4, 5}, 3))
    assert.Equal([]int{4, 4}, Maximum([]int{4, -4, 4}, 2))
    assert.Equal([]int{2}, Maximum([]int{-3, 2, 1, 2, -1, -2, 1}, 1))
  }

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMaximum(t)
}