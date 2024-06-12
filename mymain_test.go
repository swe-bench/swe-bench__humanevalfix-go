package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEvenOddCount(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([2]int{0, 1}, EvenOddCount(7))
    assert.Equal([2]int{1, 1}, EvenOddCount(-78))
    assert.Equal([2]int{2, 2}, EvenOddCount(3452))
    assert.Equal([2]int{3, 3}, EvenOddCount(346211))
    assert.Equal([2]int{3, 3}, EvenOddCount(-345821))
    assert.Equal([2]int{1, 0}, EvenOddCount(-2))
    assert.Equal([2]int{2, 3}, EvenOddCount(-45347))
    assert.Equal([2]int{1, 0}, EvenOddCount(0))
}
