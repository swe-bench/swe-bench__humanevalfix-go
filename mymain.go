package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// Given an array of integers nums, find the minimum sum of any non-empty sub-array
// of nums.
// Example
// Minsubarraysum([2, 3, 4, 1, 2, 4]) == 1
// Minsubarraysum([-1, -2, -3]) == -6
func Minsubarraysum(nums []int) int {

    max_sum := 0
    s := 0
    for _, num := range nums {
        s += -num
        if s < 0 {
            s = 0
        }
        if s > max_sum {
            max_sum = s
        }
    }
    if max_sum == 0 {
        max_sum = math.MaxInt
        for _, i := range nums {
            if -i > max_sum {
                max_sum = -i + 1
            }
        }
    }
    return -max_sum
}

func ExampleTestMinSubArraySum(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, Minsubarraysum([]int{2, 3, 4, 1, 2, 4}))
    assert.Equal(-6, Minsubarraysum([]int{-1, -2, -3}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMinsubarraysum(t)
}