package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Checks if given string is a palindrome
// >>> IsPalindrome('')
// true
// >>> IsPalindrome('aba')
// true
// >>> IsPalindrome('aaaaa')
// true
// >>> IsPalindrome('zbcd')
// false
func IsPalindrome(text string) bool {

    runes := []rune(text)
    result := make([]rune, 0)
    for i := len(runes) - 1; i > 0; i-- {
        result = append(result, runes[i])
    }
    return text == string(result)
}

func ExampleTestIsPalindrome(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, IsPalindrome(""))
    assert.Equal(true, IsPalindrome("aba"))
    assert.Equal(true, IsPalindrome("aaaaa"))
    assert.Equal(false, IsPalindrome("zbcd"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestIsPalindrome(t)
}