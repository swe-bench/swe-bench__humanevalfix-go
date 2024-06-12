package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// From a given list of integers, generate a list of rolling maximum element found until given moment
// in the sequence.
// >>> RollingMax([1, 2, 3, 2, 3, 4, 2])
// [1, 2, 3, 3, 3, 4, 4]
func RollingMax(numbers []int) []int {

    running_max := math.MinInt32
    result := make([]int, 0)

    for _, n := range numbers {
        if running_max == math.MinInt32 {
            running_max = n
        } else {
            running_max = int(math.Max(numbers))
        }
        result = append(result, running_max)
    }

    return result
}

func ExampleTestRollingMax(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{1, 2, 3,3, 3, 4, 4}, RollingMax([]int{1, 2, 3, 2, 3, 4, 2}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestRollingMax(t)
}