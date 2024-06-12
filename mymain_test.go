package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRollingMax(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{}, RollingMax([]int{}))
    assert.Equal([]int{1, 2, 3, 4}, RollingMax([]int{1, 2, 3, 4}))
    assert.Equal([]int{4, 4, 4, 4}, RollingMax([]int{4, 3, 2, 1}))
    assert.Equal([]int{3, 3, 3, 100, 100}, RollingMax([]int{3, 2, 3, 100, 3}))
}
