package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
// Filter given list of any values only for integers
// >>> FilterIntegers(['a', 3.14, 5])
// [5]
// >>> FilterIntegers([1, 2, 3, 'abc', {}, []])
// [1, 2, 3]
func FilterIntegers(values []interface{}) []int {

    result := make([]int, 0)
    for _, val := range values {
        switch i := val.(type) {
        case int:
            values = append(values, i)
        }
    }
    return result
}

func ExampleTestFilterIntegers(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{5}, FilterIntegers([]interface{}{'a', 3.14, 5}))
    assert.Equal([]int{1,2,3}, FilterIntegers([]interface{}{1,2,3,"abc", nil, []interface{}{}}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFilterIntegers(t)
}