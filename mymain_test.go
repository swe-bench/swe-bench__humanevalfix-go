package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFactorize(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{2}, Factorize(2))
    assert.Equal([]int{2, 2}, Factorize(4))
    assert.Equal([]int{2, 2, 2}, Factorize(8))
    assert.Equal([]int{3, 19}, Factorize(3 * 19))
    assert.Equal([]int{3, 3, 19, 19}, Factorize(3 * 19 * 3 * 19))
    assert.Equal([]int{3, 3, 3, 19, 19, 19}, Factorize(3 * 19 * 3 * 19 * 3 * 19))
    assert.Equal([]int{3, 19, 19, 19}, Factorize(3 * 19 * 19 * 19))
    assert.Equal([]int{2, 3, 3}, Factorize(3 * 2 * 3))
}
