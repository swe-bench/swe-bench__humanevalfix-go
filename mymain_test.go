package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTriangleArea(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(6.00, TriangleArea(3, 4, 5))
    assert.Equal(-1, TriangleArea(1, 2, 10))
    assert.Equal(8.18, TriangleArea(4, 8, 5))
    assert.Equal(1.73, TriangleArea(2, 2, 2))
    assert.Equal(-1, TriangleArea(1, 2, 3))
    assert.Equal(16.25, TriangleArea(10, 5, 7))
    assert.Equal(-1, TriangleArea(2, 6, 3))
    assert.Equal(0.43, TriangleArea(1, 1, 1))
    assert.Equal(-1, TriangleArea(2, 2, 10))
}
