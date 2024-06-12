package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIscube(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, Iscube(1))
    assert.Equal(false, Iscube(2))
    assert.Equal(true, Iscube(-1))
    assert.Equal(true, Iscube(64))
    assert.Equal(false, Iscube(180))
    assert.Equal(true, Iscube(1000))
    assert.Equal(true, Iscube(0))
    assert.Equal(false, Iscube(1729))
}
