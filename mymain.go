package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "regexp"
)

// RemoveVowels is a function that takes string and returns string without vowels.
// >>> RemoveVowels('')
// ''
// >>> RemoveVowels("abcdef\nghijklm")
// 'bcdf\nghjklm'
// >>> RemoveVowels('abcdef')
// 'bcdf'
// >>> RemoveVowels('aaaaa')
// ''
// >>> RemoveVowels('aaBAA')
// 'B'
// >>> RemoveVowels('zbcd')
// 'zbcd'
func RemoveVowels(text string) string {

    var re = regexp.MustCompile("[aeiouwyAEIOUWY]")
	text = re.ReplaceAllString(text, "")
	return text
}

func ExampleTestRemoveVowels(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("", RemoveVowels(""))
    assert.Equal("bcdf\nghjklm", RemoveVowels("abcdef\nghijklm"))
    assert.Equal("bcdf", RemoveVowels("abcdef"))
    assert.Equal("", RemoveVowels("aaaaa"))
    assert.Equal("B", RemoveVowels("aaBAA"))
    assert.Equal("zbcd", RemoveVowels("zbcd"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestRemoveVowels(t)
}