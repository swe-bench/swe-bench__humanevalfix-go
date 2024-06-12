package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFilterIntegers(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{}, FilterIntegers([]interface{}{}))
    assert.Equal([]int{4, 9}, FilterIntegers([]interface{}{4, nil, []interface{}{}, 23.2, 9, "adasd"}))
    assert.Equal([]int{3, 3, 3}, FilterIntegers([]interface{}{3, 'c', 3, 3, 'a', 'b'}))
}
