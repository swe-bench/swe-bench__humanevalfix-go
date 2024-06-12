package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestChooseNum(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(14, ChooseNum(12,15))
    assert.Equal(-1, ChooseNum(13,12))
    assert.Equal(12354, ChooseNum(33,12354))
    assert.Equal(-1, ChooseNum(5234,5233))
    assert.Equal(28, ChooseNum(6,29))
    assert.Equal(-1, ChooseNum(27,10))
    assert.Equal(-1, ChooseNum(7,7))
    assert.Equal(546, ChooseNum(546,546))
}
