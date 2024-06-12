package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMeanAbsoluteDeviation(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, math.Abs(MeanAbsoluteDeviation([]float64{1.0, 2.0, 3.0})-2.0/3.0) < 1e-6)
    assert.Equal(true, math.Abs(MeanAbsoluteDeviation([]float64{1.0, 2.0, 3.0, 4.0})-1.0) < 1e-6)
    assert.Equal(true, math.Abs(MeanAbsoluteDeviation([]float64{1.0, 2.0, 3.0, 4.0, 5.0})-6.0/5.0) < 1e-6)

}
