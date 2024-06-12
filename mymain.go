package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// Given the lengths of the three sides of a triangle. Return the area of
// the triangle rounded to 2 decimal points if the three sides form a valid triangle.
// Otherwise return -1
// Three sides make a valid triangle when the sum of any two sides is greater
// than the third side.
// Example:
// TriangleArea(3, 4, 5) == 6.00
// TriangleArea(1, 2, 10) == -1
func TriangleArea(a float64, b float64, c float64) interface{} {

    if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}
	s := (a + b + c)
	area := math.Pow(s * (s - a) * (s - b) * (s - c), 0.5)
	area = math.Round(area*100)/100
	return area
}

func ExampleTestTriangleArea(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(6.00, TriangleArea(3, 4, 5))
    assert.Equal(-1, TriangleArea(1, 2, 10))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestTriangleArea(t)
}