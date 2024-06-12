package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHasCloseElements(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, HasCloseElements([]float64{11.0, 2.0, 3.9, 4.0, 5.0, 2.2}, 0.3))
    assert.Equal(false, HasCloseElements([]float64{1.0, 2.0, 3.9, 4.0, 5.0, 2.2}, 0.05))
    assert.Equal(true, HasCloseElements([]float64{1.0, 2.0, 5.9, 4.0, 5.0}, 0.95))
    assert.Equal(false, HasCloseElements([]float64{1.0, 2.0, 5.9, 4.0, 5.0}, 0.8))
    assert.Equal(true, HasCloseElements([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 2.0}, 0.1))
    assert.Equal(true, HasCloseElements([]float64{1.1, 2.2, 3.1, 4.1, 5.1}, 1.0))
    assert.Equal(false, HasCloseElements([]float64{1.1, 2.2, 3.1, 4.1, 5.1}, 0.5))
}
