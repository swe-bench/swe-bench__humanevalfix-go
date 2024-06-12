package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{0, 2, 3, 5, 9, 123}, Unique([]int{5, 3,5, 2, 3, 3, 9, 0, 123}))
}
