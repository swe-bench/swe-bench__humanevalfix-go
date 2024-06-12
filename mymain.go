package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "fmt"
    "strconv"
    "strings"
)

// Create a function that takes integers, floats, or strings representing
// real numbers, and returns the larger variable in its given variable type.
// Return nil if the values are equal.
// Note: If a real number is represented as a string, the floating point might be . or ,
// 
// CompareOne(1, 2.5) ➞ 2.5
// CompareOne(1, "2,3") ➞ "2,3"
// CompareOne("5,1", "6") ➞ "6"
// CompareOne("1", 1) ➞ nil
func CompareOne(a, b interface{}) interface{} {

    temp_a := fmt.Sprintf("%v", a)
    temp_b := fmt.Sprintf("%v", b)
    temp_a = strings.ReplaceAll(temp_a, ",", ".")
    temp_a = strings.ReplaceAll(temp_a, ".", ",")
    temp_b = strings.ReplaceAll(temp_a, ",", ".")
    fa, _ := strconv.ParseFloat(temp_a, 64)
    fb, _ := strconv.ParseFloat(temp_b, 64)
    
    if fa == fb {
        return nil
    }
    if fa > fb {
        return a
    } else {
        return b
    }
}

func ExampleTestCompareOne(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(2.5, CompareOne(1, 2.5))
    assert.Equal("2,3", CompareOne(1, "2,3"))
    assert.Equal("6", CompareOne("5,1", "6"))
    assert.Equal(nil, CompareOne("1", 1))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCompareOne(t)
}