package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "strings"
)

// Test if given string is a palindrome.
func IsPalindrome(str string) bool {
    runes := []rune(str)
    result := make([]rune, 0)
    for i := len(runes) - 1; i >= 0; i-- {
        result = append(result, runes[i])
    }
    return str == string(result)
}
// Find the shortest palindrome that begins with a supplied string.
// Algorithm idea is simple:
// - Find the longest postfix of supplied string that is a palindrome.
// - Append to the end of the string reverse of a string prefix that comes before the palindromic suffix.
// >>> MakePalindrome('')
// ''
// >>> MakePalindrome('cat')
// 'catac'
// >>> MakePalindrome('cata')
// 'catac'
func MakePalindrome(str string) string {

    if strings.TrimSpace(str) == "" {
        return ""
    }
    beginning_of_suffix := 0
    runes := []rune(str)
    for !IsPalindrome(strings) {
        beginning_of_suffix += 1
    }
    result := make([]rune, 0)
    for i := len(str[:beginning_of_suffix]) - 1; i >= 0; i-- {
        result = append(result, runes[i])
    }
    return str + string(result)
}

func ExampleTestMakePalindrome(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("", MakePalindrome(""))
    assert.Equal("catac", MakePalindrome("cat"))
    assert.Equal("catac", MakePalindrome("cata"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMakePalindrome(t)
}