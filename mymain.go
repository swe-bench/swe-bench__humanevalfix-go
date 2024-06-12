package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "sort"
)

// Given an array of integers, sort the integers that are between 1 and 9 inclusive,
// reverse the resulting array, and then replace each digit by its corresponding name from
// "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine".
// 
// For example:
// arr = [2, 1, 1, 4, 5, 8, 2, 3]
// -> sort arr -> [1, 1, 2, 2, 3, 4, 5, 8]
// -> reverse arr -> [8, 5, 4, 3, 2, 2, 1, 1]
// return ["Eight", "Five", "Four", "Three", "Two", "Two", "One", "One"]
// 
// If the array is empty, return an empty array:
// arr = []
// return []
// 
// If the array has any strange number ignore it:
// arr = [1, -1 , 55]
// -> sort arr -> [-1, 1, 55]
// -> reverse arr -> [55, 1, -1]
// return = ['One']
func ByLength(arr []int)[]string {

    dic := map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
        4: "Four",
        5: "Five",
        6: "Six",
        7: "Seven",
        8: "Eight",
        9: "Nine",
    }
    new_arr := make([]string, 0)
    for _, item := range arr {
        if v, ok := dic[item]; ok {
            new_arr = append(new_arr, v)
        }
    }
    return new_arr
}

func ExampleTestByLength(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"Eight", "Five", "Four", "Three", "Two", "Two", "One", "One"}, ByLength([]int{2, 1, 1, 4, 5, 8, 2, 3}))
    assert.Equal([]string{}, ByLength([]int{}))
    assert.Equal([]string{"One"}, ByLength([]int{1, -1 , 55}))
    assert.Equal([]string{"Three", "Two", "One"}, ByLength([]int{1, -1, 3, 2}))
    assert.Equal([]string{"Nine", "Eight", "Four"}, ByLength([]int{9, 4, 8}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestByLength(t)
}