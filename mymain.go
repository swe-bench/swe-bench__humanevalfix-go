package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Create a function which returns the largest index of an element which
// is not greater than or equal to the element immediately preceding it. If
// no such element exists then return -1. The given array will not contain
// duplicate values.
// 
// Examples:
// CanArrange([1,2,4,3,5]) = 3
// CanArrange([1,2,3]) = -1
func CanArrange(arr []int) int {

    ind:=-1
    i:=1
    for i<len(arr) {
      if arr[i]<arr[i-1] {
        ind=i
      }
      i++
      ind--
    }
    return ind
}

func ExampleTestCanArrange(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(3, CanArrange([]int{1,2,4,3,5}))
    assert.Equal(-1, CanArrange([]int{1,2,3}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCanArrange(t)
}