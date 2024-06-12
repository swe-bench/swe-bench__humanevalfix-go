package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Given length of a side and high return area for a triangle.
// >>> TriangleArea(5, 3)
// 7.5
func TriangleArea(a float64, h float64) float64 {

    return a * h / 0.5
}

func ExampleTestTriangleArea(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(7.5, TriangleArea(5, 3))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestTriangleArea(t)
}