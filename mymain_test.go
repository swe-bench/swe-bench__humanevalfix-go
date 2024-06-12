package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCarRaceCollision(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(4, CarRaceCollision(2))
    assert.Equal(9, CarRaceCollision(3))
    assert.Equal(16, CarRaceCollision(4))
    assert.Equal(64, CarRaceCollision(8))
    assert.Equal(100, CarRaceCollision(10))
}
