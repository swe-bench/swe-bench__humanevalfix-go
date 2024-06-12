package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMatchParens(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("Yes", MatchParens([]string{"()(", ")"}))
    assert.Equal("No", MatchParens([]string{")", ")"}))
    assert.Equal("No", MatchParens([]string{"(()(())", "())())"}))
    assert.Equal("Yes", MatchParens([]string{")())", "(()()("}))
    assert.Equal("Yes", MatchParens([]string{"(())))", "(()())(("}))
    assert.Equal("No", MatchParens([]string{"()", "())"}))
    assert.Equal("Yes", MatchParens([]string{"(()(", "()))()"}))
    assert.Equal("No", MatchParens([]string{"((((", "((())"}))
    assert.Equal("No", MatchParens([]string{")(()", "(()("}))
    assert.Equal("No", MatchParens([]string{")(", ")("}))
    assert.Equal("Yes", MatchParens([]string{"(", ")"}))
    assert.Equal("Yes", MatchParens([]string{")", "("}))
}
