package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFindClosestElements(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([2]float64{3.9, 4.0}, FindClosestElements([]float64{1.0, 2.0, 3.9, 4.0, 5.0, 2.2}))
    assert.Equal([2]float64{5.0, 5.9}, FindClosestElements([]float64{1.0, 2.0, 5.9, 4.0, 5.0}))
    assert.Equal([2]float64{2.0, 2.2}, FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.2}))
    assert.Equal([2]float64{2.0, 2.0}, FindClosestElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0}))
    assert.Equal([2]float64{2.2, 3.1}, FindClosestElements([]float64{1.1, 2.2, 3.1, 4.1, 5.1}))
}
