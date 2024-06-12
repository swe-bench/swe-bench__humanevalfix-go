package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHistogram(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(map[rune]int{'a':2,'b': 2}, Histogram("a b b a"))
    assert.Equal(map[rune]int{'a': 2, 'b': 2}, Histogram("a b c a b"))
    assert.Equal(map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'g': 1}, Histogram("a b c d g"))
    assert.Equal(map[rune]int{'r': 1,'t': 1,'g': 1}, Histogram("r t g"))
    assert.Equal(map[rune]int{'b': 4}, Histogram("b b b b a"))
    assert.Equal(map[rune]int{}, Histogram(""))
    assert.Equal(map[rune]int{'a': 1}, Histogram("a"))
}
