package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPairsSumToZero(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, PairsSumToZero([]int{1, 3, 5, 0}))
    assert.Equal(false, PairsSumToZero([]int{1, 3, -2, 1}))
    assert.Equal(false, PairsSumToZero([]int{1, 2, 3, 7}))
    assert.Equal(true, PairsSumToZero([]int{2, 4, -5, 3, 5, 7}))
    assert.Equal(false, PairsSumToZero([]int{1}))
    assert.Equal(true, PairsSumToZero([]int{-3, 9, -1, 3, 2, 30}))
    assert.Equal(true, PairsSumToZero([]int{-3, 9, -1, 3, 2, 31}))
    assert.Equal(false, PairsSumToZero([]int{-3, 9, -1, 4, 2, 30}))
    assert.Equal(false, PairsSumToZero([]int{-3, 9, -1, 4, 2, 31}))
}
