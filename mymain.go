package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strconv"
)

// Given a non-empty array of integers arr and an integer k, return
// the sum of the elements with at most two digits from the first k elements of arr.
// 
// Example:
// 
// Input: arr = [111,21,3,4000,5,6,7,8,9], k = 4
// Output: 24 # sum of 21 + 3
// 
// Constraints:
// 1. 1 <= len(arr) <= 100
// 2. 1 <= k <= len(arr)
func AddElements(arr []int, k int) int {

    sum := 0
    for _, elem := range arr {
        if len(strconv.Itoa(elem)) <= 2 {
            sum += elem
        }
    }
    return sum
}

func ExampleTestAddElements(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(24, AddElements([]int{111,21,3,4000,5,6,7,8,9}, 4))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestAddElements(t)
}