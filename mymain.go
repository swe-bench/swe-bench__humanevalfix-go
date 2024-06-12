package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Given a non-empty list of integers lst. Add the even elements that are at odd indices..
// 
// Examples:
// Add([4, 2, 6, 7]) ==> 2
func Add(lst []int) int {

    sum := 0
    for i := 1; i < len(lst); i += 1 {
        if lst[i]%2 == 0 {
            sum += lst[i]
        }
    }
    return sum
}

func ExampleTestAdd(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(2, Add([]int{4, 2, 6, 7}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestAdd(t)
}