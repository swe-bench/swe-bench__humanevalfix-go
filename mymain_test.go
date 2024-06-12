package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([][2]int{{0, 0}, {1, 4}, {1, 0}, {2, 5}, {2, 0}}, GetRow([][]int{
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 1, 6},
        {1, 2, 3, 4, 5, 1},
    }, 1))
    assert.Equal([][2]int{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}}, GetRow([][]int{
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 5, 6},
    }, 2))
    assert.Equal([][2]int{{0, 0}, {1, 0}, {2, 1}, {2, 0}, {3, 2}, {3, 0}, {4, 3}, {4, 0}, {5, 4}, {5, 0}, {6, 5}, {6, 0}}, GetRow([][]int{
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 5, 6},
        {1, 1, 3, 4, 5, 6},
        {1, 2, 1, 4, 5, 6},
        {1, 2, 3, 1, 5, 6},
        {1, 2, 3, 4, 1, 6},
        {1, 2, 3, 4, 5, 1},
    }, 1))
    assert.Equal([][2]int{}, GetRow([][]int{{}}, 1))
    assert.Equal([][2]int{}, GetRow([][]int{{1}}, 2))
    assert.Equal([][2]int{{2, 2}}, GetRow([][]int{{}, {1}, {1, 2, 3}}, 3))
}
