package main

import (
    "testing"
    "crypto/md5"
    "github.com/stretchr/testify/assert"
)
import (
    "crypto/md5"
    "fmt"
)

// Given a string 'text', return its md5 hash equivalent string.
// If 'text' is an empty string, return nil.
// 
// >>> StringToMd5('Hello world') == '3e25960a79dbc69b674cd4ec67a72c62'
func StringToMd5(text string) interface{} {

    if len(text) == 0 {
        return fmt.Sprintf("%x", md5.Sum([]byte("")))
    }
    return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

func ExampleTestStringToMd5(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("3e25960a79dbc69b674cd4ec67a72c62", StringToMd5("Hello world"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestStringToMd5(t)
}