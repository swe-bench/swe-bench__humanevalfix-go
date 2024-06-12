package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsBored(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, IsBored("Hello world"), "Test 1")
    assert.Equal(0, IsBored("Is the sky blue?"), "Test 2")
    assert.Equal(1, IsBored("I love It !"), "Test 3")
    assert.Equal(0, IsBored("bIt"), "Test 4")
    assert.Equal(2, IsBored("I feel good today. I will be productive. will kill It"), "Test 5")
    assert.Equal(0, IsBored("You and I are going for a walk"), "Test 6")
}
