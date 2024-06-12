package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// Write a function that takes an integer a and returns true
// if this ingeger is a cube of some integer number.
// Note: you may assume the input is always valid.
// Examples:
// Iscube(1) ==> true
// Iscube(2) ==> false
// Iscube(-1) ==> true
// Iscube(64) ==> true
// Iscube(0) ==> true
// Iscube(180) ==> false
func Iscube(a int) bool {

    abs := math.Abs(float64(a))
	return int(math.Round(math.Pow(abs, 1.0/3.0))) == int(abs)
}

func ExampleTestIscube(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, Iscube(1))
    assert.Equal(false, Iscube(2))
    assert.Equal(true, Iscube(-1))
    assert.Equal(true, Iscube(64))
    assert.Equal(false, Iscube(180))
    assert.Equal(true, Iscube(0))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestIscube(t)
}