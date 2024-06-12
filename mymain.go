package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Task
// Write a function that takes a string as input and returns the sum of the upper characters only'
// ASCII codes.
// 
// Examples:
// Digitsum("") => 0
// Digitsum("abAB") => 131
// Digitsum("abcCd") => 67
// Digitsum("helloE") => 69
// Digitsum("woArBld") => 131
// Digitsum("aAaaaXa") => 153
func Digitsum(x string) int {

    if len(x) == 0 {
		return 0
	}
	result := 0
	for _, i := range x {
		if 'a' <= i && i <= 'z' {
			result += int(i)
		}
	}
	return result
}

func ExampleTestDigitSum(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, Digitsum(""))
    assert.Equal(131, Digitsum("abAB"))
    assert.Equal(67, Digitsum("abcCd"))
    assert.Equal(69, Digitsum("helloE"))
    assert.Equal(131, Digitsum("woArBld"))
    assert.Equal(153, Digitsum("aAaaaXa"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestDigitsum(t)
}