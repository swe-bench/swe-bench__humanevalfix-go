package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// You are given a rectangular grid of wells. Each row represents a single well,
// and each 1 in a row represents a single unit of water.
// Each well has a corresponding bucket that can be used to extract water from it,
// and all buckets have the same capacity.
// Your task is to use the buckets to empty the wells.
// Output the number of times you need to lower the buckets.
// 
// Example 1:
// Input:
// grid : [[0,0,1,0], [0,1,0,0], [1,1,1,1]]
// bucket_capacity : 1
// Output: 6
// 
// Example 2:
// Input:
// grid : [[0,0,1,1], [0,0,0,0], [1,1,1,1], [0,1,1,1]]
// bucket_capacity : 2
// Output: 5
// 
// Example 3:
// Input:
// grid : [[0,0,0], [0,0,0]]
// bucket_capacity : 5
// Output: 0
// 
// Constraints:
// * all wells have the same length
// * 1 <= grid.length <= 10^2
// * 1 <= grid[:,1].length <= 10^2
// * grid[i][j] -> 0 | 1
// * 1 <= capacity <= 10
func MaxFill(grid [][]int, capacity int) int {

    result := 0
    for _, arr := range grid {
        sum := 0
        for _, i := range arr {
            sum += i
        }
        result += int(math.Floor(float64(sum) / float64(capacity)))
    }
    return result
}

func ExampleTestMaxFill(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(6, MaxFill([][]int{{0,0,1,0}, {0,1,0,0}, {1,1,1,1}}, 1))
  assert.Equal(5, MaxFill([][]int{{0,0,1,1}, {0,0,0,0}, {1,1,1,1}, {0,1,1,1}}, 2))
  assert.Equal(0, MaxFill([][]int{{0,0,0}, {0,0,0}}, 5))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMaxFill(t)
}