package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCheckDictCase(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, CheckDictCase(map[interface{}]interface{}{"p": "pineapple", "b": "banana"}))
    assert.Equal(false, CheckDictCase(map[interface{}]interface{}{"p": "pineapple", "A": "banana", "B": "banana"}))
    assert.Equal(false, CheckDictCase(map[interface{}]interface{}{"p": "pineapple", 5: "banana", "a": "apple"}))
    assert.Equal(false, CheckDictCase(map[interface{}]interface{}{"Name": "John", "Age": "36", "City": "Houston"}))
    assert.Equal(true, CheckDictCase(map[interface{}]interface{}{"STATE": "NC", "ZIP": "12345"}))
    assert.Equal(true, CheckDictCase(map[interface{}]interface{}{"fruit": "Orange", "taste": "Sweet"}))
    assert.Equal(false, CheckDictCase(map[interface{}]interface{}{}))
}
