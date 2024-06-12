package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSkjkasdkd(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(10, Skjkasdkd([]int{0, 3, 2, 1, 3, 5, 7, 4, 5, 5, 5, 2, 181, 32, 4, 32, 3, 2, 32, 324, 4, 3}))
    assert.Equal(25, Skjkasdkd([]int{1, 0, 1, 8, 2, 4597, 2, 1, 3, 40, 1, 2, 1, 2, 4, 2, 5, 1}))
    assert.Equal(13, Skjkasdkd([]int{1, 3, 1, 32, 5107, 34, 83278, 109, 163, 23, 2323, 32, 30, 1, 9, 3}))
    assert.Equal(11, Skjkasdkd([]int{0, 724, 32, 71, 99, 32, 6, 0, 5, 91, 83, 0, 5, 6}))
    assert.Equal(3, Skjkasdkd([]int{0, 81, 12, 3, 1, 21}))
    assert.Equal(7, Skjkasdkd([]int{0, 8, 1, 2, 1, 7}))
    assert.Equal(19, Skjkasdkd([]int{8191}))
    assert.Equal(19, Skjkasdkd([]int{8191, 123456, 127, 7}))
    assert.Equal(10, Skjkasdkd([]int{127, 97, 8192}))
}
