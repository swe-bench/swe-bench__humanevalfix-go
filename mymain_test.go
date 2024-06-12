package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTri(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]float64{1, 3, 2.0, 8.0}, Tri(3))
    assert.Equal([]float64{1, 3, 2.0, 8.0, 3.0}, Tri(4))
    assert.Equal([]float64{1, 3, 2.0, 8.0, 3.0, 15.0}, Tri(5))
    assert.Equal([]float64{1, 3, 2.0, 8.0, 3.0, 15.0, 4.0}, Tri(6))
    assert.Equal([]float64{1, 3, 2.0, 8.0, 3.0, 15.0, 4.0, 24.0}, Tri(7))
    assert.Equal([]float64{1, 3, 2.0, 8.0, 3.0, 15.0, 4.0, 24.0, 5.0}, Tri(8))
    assert.Equal([]float64{1, 3, 2.0, 8.0, 3.0, 15.0, 4.0, 24.0, 5.0, 35.0}, Tri(9))
    assert.Equal([]float64{1, 3, 2.0, 8.0, 3.0, 15.0, 4.0, 24.0, 5.0, 35.0, 6.0, 48.0, 7.0, 63.0, 8.0, 80.0, 9.0, 99.0, 10.0, 120.0, 11.0}, Tri(20))
    assert.Equal([]float64{1}, Tri(0))
    assert.Equal([]float64{1, 3}, Tri(1))
}
