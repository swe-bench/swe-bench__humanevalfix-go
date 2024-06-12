package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestWordsString(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"Hi", "my", "name", "is", "John"}, WordsString("Hi, my name is John"))
    assert.Equal([]string{"One", "two", "three", "four", "five", "six"}, WordsString("One, two, three, four, five, six"))
    assert.Equal([]string{}, WordsString(""))
    assert.Equal([]string{"Hi", "my", "name"}, WordsString("Hi, my name"))
    assert.Equal([]string{"One", "two", "three", "four", "five", "six"}, WordsString("One,, two, three, four, five, six,"))
    assert.Equal([]string{"ahmed", "gamal"}, WordsString("ahmed     , gamal"))
}
