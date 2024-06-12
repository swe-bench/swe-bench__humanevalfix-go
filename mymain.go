package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Insert a number 'delimeter' between every two consecutive elements of input list `numbers'
// >>> Intersperse([], 4)
// []
// >>> Intersperse([1, 2, 3], 4)
// [1, 4, 2, 4, 3]
func Intersperse(numbers []int, delimeter int) []int {

    result := make([]int, 0)
    if len(numbers) == 0 {
        return result
    }
    for i := 0; i < len(numbers)-1; i++ {
        n := numbers[i]
        result = append(result, n)
        result = append(result, delimeter)
    }
    return result
}

func ExampleTestIntersperse(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{}, Intersperse([]int{}, 4))
    assert.Equal([]int{1,4,2,4,3}, Intersperse([]int{1,2,3}, 4))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestIntersperse(t)
}