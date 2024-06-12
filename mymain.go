package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// This function takes two positive numbers x and y and returns the
// biggest even integer number that is in the range [x, y] inclusive. If
// there's no such number, then the function should return -1.
// 
// For example:
// ChooseNum(12, 15) = 14
// ChooseNum(13, 12) = -1
func ChooseNum(x, y int) int {

    if x > y {
        return -1
    }
    if y % 2 == 0 {
        return y
    }
    if x == y {
        return -1
    }
    return x - 1
}

func ExampleTestChooseNum(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(14, ChooseNum(12,15))
    assert.Equal(-1, ChooseNum(13,12))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestChooseNum(t)
}