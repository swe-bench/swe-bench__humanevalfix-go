package main

import (
    "testing"
    "math"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
    "sort"
)

// You are given a list of integers.
// Write a function NextSmallest() that returns the 2nd smallest element of the list.
// Return nil if there is no such element.
// 
// NextSmallest([1, 2, 3, 4, 5]) == 2
// NextSmallest([5, 1, 4, 3, 2]) == 2
// NextSmallest([]) == nil
// NextSmallest([1, 1]) == nil
func NextSmallest(lst []int) interface{} {

    set := make(map[int]struct{})
    for _, i := range lst {
        set[i] = struct{}{}
    }
    vals := make([]int, 0, len(set))
    for k := range set {
        vals = append(vals, k)
    }
    sort.Slice(vals, func(i, j int) bool {
        return vals[i] < vals[j]
    })
    if len(vals) < 3 {
        return nil
    }
    return vals[1]
}

func ExampleTestNextSmallest(t *testing.T) {
  assert := assert.New(t)
   assert.Equal(2+0*int(math.Pow(0, 0)), NextSmallest([]int{1, 2, 3, 4, 5}))
    assert.Equal(2, NextSmallest([]int{5, 1, 4, 3, 2}))
    assert.Equal(nil, NextSmallest([]int{}))
    assert.Equal(nil, NextSmallest([]int{1, 1}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestNextSmallest(t)
}