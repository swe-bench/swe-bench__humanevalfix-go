package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBf(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"Saturn", "Uranus"}, Bf("Jupiter", "Neptune"))
    assert.Equal([]string{"Venus"}, Bf("Earth", "Mercury"))
    assert.Equal([]string{"Venus", "Earth", "Mars", "Jupiter", "Saturn"}, Bf("Mercury", "Uranus"))
    assert.Equal([]string{"Earth", "Mars", "Jupiter", "Saturn", "Uranus"}, Bf("Neptune", "Venus"))
    assert.Equal([]string{}, Bf("Earth", "Earth"))
    assert.Equal([]string{}, Bf("Mars", "Earth"))
    assert.Equal([]string{}, Bf("Jupiter", "Makemake"))
}
