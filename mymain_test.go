package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTriplesSumToZero(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, TriplesSumToZero([]int{1, 3, 5, 0}))
    assert.Equal(false, TriplesSumToZero([]int{1, 3, 5, -1}))
    assert.Equal(true, TriplesSumToZero([]int{1, 3, -2, 1}))
    assert.Equal(false, TriplesSumToZero([]int{1, 2, 3, 7}))
    assert.Equal(false, TriplesSumToZero([]int{1, 2, 5, 7}))
    assert.Equal(true, TriplesSumToZero([]int{2, 4, -5, 3, 9, 7}))
    assert.Equal(false, TriplesSumToZero([]int{1}))
    assert.Equal(false, TriplesSumToZero([]int{1, 3, 5, -100}))
    assert.Equal(false, TriplesSumToZero([]int{100, 3, 5, -100}))
}
