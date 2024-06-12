package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsNested(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, IsNested("[[]]"))
    assert.Equal(false, IsNested("[]]]]]]][[[[[]"))
    assert.Equal(false, IsNested("[][]"))
    assert.Equal(false, IsNested("'[]'"))
    assert.Equal(true, IsNested("[[[[]]]]"))
    assert.Equal(false, IsNested("[]]]]]]]]]]"))
    assert.Equal(true, IsNested("[][][[]]"))
    assert.Equal(false, IsNested("[[]"))
    assert.Equal(false, IsNested("[]]"))
    assert.Equal(true, IsNested("[[]][["))
    assert.Equal(true, IsNested("[[][]]"))
    assert.Equal(false, IsNested(""))
    assert.Equal(false, IsNested("[[[[[[[["))
    assert.Equal(false, IsNested("]]]]]]]]"))
}
