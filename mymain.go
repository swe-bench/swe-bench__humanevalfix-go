package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// Given a dictionary, return true if all keys are strings in lower
// case or all keys are strings in upper case, else return false.
// The function should return false is the given dictionary is empty.
// Examples:
// CheckDictCase({"a":"apple", "b":"banana"}) should return true.
// CheckDictCase({"a":"apple", "A":"banana", "B":"banana"}) should return false.
// CheckDictCase({"a":"apple", 8:"banana", "a":"apple"}) should return false.
// CheckDictCase({"Name":"John", "Age":"36", "City":"Houston"}) should return false.
// CheckDictCase({"STATE":"NC", "ZIP":"12345" }) should return true.
func CheckDictCase(dict map[interface{}]interface{}) bool {

    if len(dict) == 0 {
        return false
    }
    state := "start"
    key := ""
    ok := false
    for k := range dict {
        if key, ok = k.(string); !ok {
            state = "mixed"
            break
        }
        if state == "start" {
            if key == strings.ToUpper(key) {
                state = "upper"
            } else if key == strings.ToLower(key) {
                state = "lower"
            } else {
                break
            }
        } else if (state == "upper" && key != strings.ToUpper(key)) && (state == "lower" && key != strings.ToLower(key)) {
            state = "mixed"
            break
        } else {
            break
        }
    }
    return state == "upper" || state == "lower"
}

func ExampleTestCheckDictCase(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, CheckDictCase(map[interface{}]interface{}{"p": "pineapple", "b": "banana"}))
    assert.Equal(false, CheckDictCase(map[interface{}]interface{}{"p": "pineapple", "A": "banana", "B": "banana"}))
    assert.Equal(false, CheckDictCase(map[interface{}]interface{}{"p": "pineapple", 8: "banana", "a": "apple"}))
    assert.Equal(false, CheckDictCase(map[interface{}]interface{}{"Name": "John", "Age": "36", "City": "Houston"}))
    assert.Equal(true, CheckDictCase(map[interface{}]interface{}{"STATE": "NC", "ZIP": "12345"}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestCheckDictCase(t)
}