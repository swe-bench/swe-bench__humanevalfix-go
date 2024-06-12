package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Given the lengths of the three sides of a triangle. Return true if the three
// sides form a right-angled triangle, false otherwise.
// A right-angled triangle is a triangle in which one angle is right angle or
// 90 degree.
// Example:
// RightAngleTriangle(3, 4, 5) == true
// RightAngleTriangle(1, 2, 3) == false
func RightAngleTriangle(a, b, c int) bool {

    return c*c == a*a + b*b
}

func ExampleTestRightAngleTriangle(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, RightAngleTriangle(3, 4, 5))
    assert.Equal(false, RightAngleTriangle(1, 2, 3))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestRightAngleTriangle(t)
}