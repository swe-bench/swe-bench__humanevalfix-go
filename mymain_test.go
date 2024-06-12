package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetMaxTriples(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, GetMaxTriples(5))
    assert.Equal(4, GetMaxTriples(6))
    assert.Equal(36, GetMaxTriples(10))
    assert.Equal(53361, GetMaxTriples(100))
} 
