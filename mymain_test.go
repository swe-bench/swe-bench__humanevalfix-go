package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDoAlgebra(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(37, DoAlgebra([]string{"**", "*", "+"}, []int{2, 3, 4, 5}))
    assert.Equal(9, DoAlgebra([]string{"+", "*", "-"}, []int{2, 3, 4, 5}))
    assert.Equal(8, DoAlgebra([]string{"//", "*"}, []int{7, 3, 4}))
}
