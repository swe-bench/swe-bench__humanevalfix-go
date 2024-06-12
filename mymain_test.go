package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParseNestedParens(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{2, 3, 1, 3}, ParseNestedParens("(()()) ((())) () ((())()())"))
    assert.Equal([]int{1, 2, 3, 4}, ParseNestedParens("() (()) ((())) (((())))"))
    assert.Equal([]int{4}, ParseNestedParens("(()(())((())))"))
}
