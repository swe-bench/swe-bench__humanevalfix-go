package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMoveOneBall(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, MoveOneBall([]int{3, 4, 5, 1, 2}))
    assert.Equal(true, MoveOneBall([]int{3, 5, 10, 1, 2}))
    assert.Equal(false, MoveOneBall([]int{4, 3, 1, 2}))
    assert.Equal(false, MoveOneBall([]int{3, 5, 4, 1, 2}))
    assert.Equal(true, MoveOneBall([]int{}))
}
