package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTriangleArea(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(7.5, TriangleArea(5, 3))
    assert.Equal(2.0, TriangleArea(2, 2))
    assert.Equal(40.0, TriangleArea(10, 8))
}
