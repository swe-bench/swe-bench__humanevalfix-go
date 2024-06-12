package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParseMusic(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{}, ParseMusic(""))
    assert.Equal([]int{4, 4, 4, 4}, ParseMusic("o o o o"))
    assert.Equal([]int{1, 1, 1, 1}, ParseMusic(".| .| .| .|"))
    assert.Equal([]int{2, 2, 1, 1, 4, 4, 4, 4}, ParseMusic("o| o| .| .| o o o o"))
    assert.Equal([]int{2, 1, 2, 1, 4, 2, 4, 2}, ParseMusic("o| .| o| .| o o| o o|"))
}
