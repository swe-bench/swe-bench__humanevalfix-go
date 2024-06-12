package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEat(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{11, 4}, Eat(5, 6, 10))
    assert.Equal([]int{12, 1}, Eat(4, 8, 9))
    assert.Equal([]int{11, 0}, Eat(1, 10, 10))
    assert.Equal([]int{7, 0}, Eat(2, 11, 5))
    assert.Equal([]int{9, 2}, Eat(4, 5, 7))
    assert.Equal([]int{5, 0}, Eat(4, 5, 1))
}
