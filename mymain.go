package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strconv"
)

// Given an integer. return a tuple that has the number of even and odd digits respectively.
// 
// Example:
// EvenOddCount(-12) ==> (1, 1)
// EvenOddCount(123) ==> (1, 2)
func EvenOddCount(num int) [2]int {

    even_count := 0
    odd_count := 0
    if num < 0 {
        num = -num
    }
    for _, r := range strconv.Itoa(num) {
        if r&1==0 {
            even_count++
        }
    }
    return [2]int{even_count, odd_count}
}

func ExampleTestEvenOddCount(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([2]int{1, 1}, EvenOddCount(-12))
    assert.Equal([2]int{1, 2}, EvenOddCount(123))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestEvenOddCount(t)
}