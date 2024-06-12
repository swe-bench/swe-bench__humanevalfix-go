package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
    "sort"
)

// We have an array 'arr' of N integers arr[1], arr[2], ..., arr[N].The
// numbers in the array will be randomly ordered. Your task is to determine if
// it is possible to get an array sorted in non-decreasing order by performing
// the following operation on the given array:
// You are allowed to perform right shift operation any number of times.
// 
// One right shift operation means shifting all elements of the array by one
// position in the right direction. The last element of the array will be moved to
// the starting position in the array i.e. 0th index.
// 
// If it is possible to obtain the sorted array by performing the above operation
// then return true else return false.
// If the given array is empty then return true.
// 
// Note: The given list is guaranteed to have unique elements.
// 
// For Example:
// 
// MoveOneBall([3, 4, 5, 1, 2])==>true
// Explanation: By performin 2 right shift operations, non-decreasing order can
// be achieved for the given array.
// MoveOneBall([3, 5, 4, 1, 2])==>false
// Explanation:It is not possible to get non-decreasing order for the given
// array by performing any number of right shift operations.
func MoveOneBall(arr []int) bool {

    if len(arr)==0 {
      return true
    }
    sorted_array := make([]int, len(arr))
    copy(sorted_array, arr)
    sort.Slice(sorted_array, func(i, j int) bool {
        return sorted_array[i] < sorted_array[j]
    })    
    min_value := math.MaxInt
    min_index := -1
    for i, x := range arr {
        if i < min_value {
            min_index, min_value = i, x
        }
    }
    my_arr := make([]int, len(arr[min_index:]))
    copy(my_arr, arr[min_index:])
    my_arr = append(my_arr, arr[0:min_index]...)
    for i :=0;i<len(arr);i++ {
      if my_arr[i]!=arr[i]{
        return false
      }
    }
    return true
}

func ExampleTestMoveOneBall(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, MoveOneBall([]int{3, 4, 5, 1, 2}))
    assert.Equal(false, MoveOneBall([]int{3, 5, 4, 1, 2}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMoveOneBall(t)
}