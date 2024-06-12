package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, IsHappy("a"), "a")
    assert.Equal(false, IsHappy("aa"), "aa")
    assert.Equal(true, IsHappy("abcd"), "abcd")
    assert.Equal(false, IsHappy("aabb"), "aabb")
    assert.Equal(true, IsHappy("adb"), "adb")
    assert.Equal(false, IsHappy("xyy"), "xyy")
    assert.Equal(true, IsHappy("iopaxpoi"), "iopaxpoi")
    assert.Equal(false, IsHappy("iopaxioi"), "iopaxioi")
}
