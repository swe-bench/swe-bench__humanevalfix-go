package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTotalMatch(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{}, TotalMatch([]string{}, []string{}))
    assert.Equal([]string{"hi", "hi"}, TotalMatch([]string{"hi", "admin"}, []string{"hi", "hi"}))
    assert.Equal([]string{"hi", "admin"}, TotalMatch([]string{"hi", "admin"}, []string{"hi", "hi", "admin", "project"}))
    assert.Equal([]string{"4"}, TotalMatch([]string{"4"},[]string{"1", "2", "3", "4", "5"}))
    assert.Equal([]string{"hI", "Hi"}, TotalMatch([]string{"hi", "admin"}, []string{"hI", "Hi"}))
    assert.Equal([]string{"hI", "hi", "hi"}, TotalMatch([]string{"hi", "admin"}, []string{"hI", "hi", "hi"}))
    assert.Equal([]string{"hi", "admin"}, TotalMatch([]string{"hi", "admin"}, []string{"hI", "hi", "hii"}))
    assert.Equal([]string{}, TotalMatch([]string{}, []string{"this"}))
    assert.Equal([]string{}, TotalMatch([]string{"this"}, []string{}))
}
