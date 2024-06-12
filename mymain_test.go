package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSpecialFactorial(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(288, SpecialFactorial(4))
    assert.Equal(34560, SpecialFactorial(5))
    assert.Equal(125411328000, SpecialFactorial(7))
    assert.Equal(1, SpecialFactorial(1))
}
