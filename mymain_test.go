package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestByLength(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"Eight", "Five", "Four", "Three", "Two", "Two", "One", "One"}, ByLength([]int{2, 1, 1, 4, 5, 8, 2, 3}))
    assert.Equal([]string{}, ByLength([]int{}))
    assert.Equal([]string{"One"}, ByLength([]int{1, -1 , 55}))
    assert.Equal([]string{"Three", "Two", "One"}, ByLength([]int{1, -1, 3, 2}))
    assert.Equal([]string{"Nine", "Eight", "Four"}, ByLength([]int{9, 4, 8}))
}
