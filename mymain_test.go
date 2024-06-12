package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSortedListSum(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"aa"}, SortedListSum([]string{"aa", "a", "aaa"}))
    assert.Equal([]string{"AI", "asdf", "school"}, SortedListSum([]string{"school", "AI", "asdf", "b"}))
    assert.Equal([]string{}, SortedListSum([]string{"d", "b", "c", "a"}))
    assert.Equal([]string{"abcd", "dcba"}, SortedListSum([]string{"d", "dcba", "abcd", "a"}))
    assert.Equal([]string{"AI", "ai", "au"}, SortedListSum([]string{"AI", "ai", "au"}))
    assert.Equal([]string{}, SortedListSum([]string{"a", "b", "b", "c", "c", "a"}))
    assert.Equal([]string{"cc", "dd", "aaaa", "bbbb"}, SortedListSum([]string{"aaaa", "bbbb", "dd", "cc"}))
}
