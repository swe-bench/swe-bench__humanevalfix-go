package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Find how many times a given substring can be found in the original string. Count overlaping cases.
// >>> HowManyTimes('', 'a')
// 0
// >>> HowManyTimes('aaa', 'a')
// 3
// >>> HowManyTimes('aaaa', 'aa')
// 3
func HowManyTimes(str string,substring string) int{

    times := 0
	for i := 0; i < (len(str) - len(substring)); i++ {
		if str[i:i+len(substring)] == substring {
			times += 1
		}
	}
	return times
}


func ExampleTestHowManyTimes(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, HowManyTimes("", "a"))
    assert.Equal(3, HowManyTimes("aaa", "a"))
    assert.Equal(3, HowManyTimes("aaaa", "aa"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestHowManyTimes(t)
}