package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
    "strconv"
)

// You are given a list of integers.
// You need to find the largest prime value and return the sum of its digits.
// 
// Examples:
// For lst = [0,3,2,1,3,5,7,4,5,5,5,2,181,32,4,32,3,2,32,324,4,3] the output should be 10
// For lst = [1,0,1,8,2,4597,2,1,3,40,1,2,1,2,4,2,5,1] the output should be 25
// For lst = [1,3,1,32,5107,34,83278,109,163,23,2323,32,30,1,9,3] the output should be 13
// For lst = [0,724,32,71,99,32,6,0,5,91,83,0,5,6] the output should be 11
// For lst = [0,81,12,3,1,21] the output should be 3
// For lst = [0,8,1,2,1,7] the output should be 7
func Skjkasdkd(lst []int) int {

    isPrime := func(n int) bool {
        for i := 2; i < int(math.Pow(float64(n), 0.5)+1); i++ {
            if n%i == 0 {
                return true
            }
        }
        return false
    }
    maxx := 0
    i := 0
    for i < len(lst) {
        if lst[i] > maxx && isPrime(lst[i]) {
            maxx = lst[i]
        }
        i++
    }
    sum := 0
    for _, d := range strconv.Itoa(maxx) {
        sum += int(d - '0')
    }
    return sum
}

func ExampleTestSkjkasdkd(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(10, Skjkasdkd([]int{0, 3, 2, 1, 3, 5, 7, 4, 5, 5, 5, 2, 181, 32, 4, 32, 3, 2, 32, 324, 4, 3}))
    assert.Equal(25, Skjkasdkd([]int{1, 0, 1, 8, 2, 4597, 2, 1, 3, 40, 1, 2, 1, 2, 4, 2, 5, 1}))
    assert.Equal(13, Skjkasdkd([]int{1, 3, 1, 32, 5107, 34, 83278, 109, 163, 23, 2323, 32, 30, 1, 9, 3}))
    assert.Equal(11, Skjkasdkd([]int{0, 724, 32, 71, 99, 32, 6, 0, 5, 91, 83, 0, 5, 6}))
    assert.Equal(3, Skjkasdkd([]int{0, 81, 12, 3, 1, 21}))
    assert.Equal(7, Skjkasdkd([]int{0, 8, 1, 2, 1, 7}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSkjkasdkd(t)
}