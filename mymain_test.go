package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSeparateParenGroups(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"(()())", "((()))", "()", "((())()())"}, SeparateParenGroups("(()()) ((())) () ((())()())"))
    assert.Equal([]string{"()", "(())", "((()))", "(((())))"}, SeparateParenGroups("() (()) ((())) (((())))"))
    assert.Equal([]string{"(()(())((())))"}, SeparateParenGroups("(()(())((())))"))
    assert.Equal([]string{"()", "(())", "(()())"}, SeparateParenGroups("( ) (( )) (( )( ))"))
}
