package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Given a non-empty list of integers, return the sum of all of the odd elements that are in even positions.
// 
// Examples
// Solution([5, 8, 7, 1]) ==> 12
// Solution([3, 3, 3, 3, 3]) ==> 9
// Solution([30, 13, 24, 321]) ==>0
func Solution(lst []int) int {

    sum:=0
    for i, x := range lst {
        if i&1==1&&x&1==1 {
            sum+=x
        }
    }
    return sum
}

func ExampleTestSolution(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(12, Solution([]int{5, 8, 7, 1}))
    assert.Equal(9, Solution([]int{3, 3, 3, 3, 3}))
    assert.Equal(0, Solution([]int{30, 13, 24, 321}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSolution(t)
}