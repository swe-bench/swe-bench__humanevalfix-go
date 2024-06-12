package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Circular shift the digits of the integer x, shift the digits right by shift
// and return the result as a string.
// If shift > number of digits, return digits reversed.
// >>> CircularShift(12, 1)
// "21"
// >>> CircularShift(12, 2)
// "12"
func CircularShift(x int,shift int) string {

    s := strconv.Itoa(x)
	if shift > len(s) {
		runes := make([]rune, 0)
		for i := len(s)-1; i >= 0; i-- {
			runes = append(runes, rune(s[i]))
		}
		return string(runes)
	}else{
		return s[:len(s)-shift]+s[len(s)-shift:]
	}
}

func ExampleTestCircularShift(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("12", CircularShift(12, 2))
    assert.Equal("21", CircularShift(12, 1))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCircularShift(t)
}