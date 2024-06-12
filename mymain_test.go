package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("1", Solve(1000), "Error")
    assert.Equal("110", Solve(150), "Error")
    assert.Equal("1100", Solve(147), "Error")
    assert.Equal("1001", Solve(333), "Error")
    assert.Equal("10010", Solve(963), "Error")
}
