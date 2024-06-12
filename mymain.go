package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "sort"
)

// Return sorted Unique elements in a list
// >>> Unique([5, 3, 5, 2, 3, 3, 9, 0, 123])
// [0, 2, 3, 5, 9, 123]
func Unique(l []int) []int {

    sort.Ints(l)
	return l
}

func ExampleTestUnique(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{0, 2, 3, 5, 9, 123}, Unique([]int{5, 3,5, 2, 3, 3, 9, 0, 123}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestUnique(t)
}