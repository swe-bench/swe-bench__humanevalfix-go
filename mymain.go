package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Input to this function is a string containing multiple groups of nested parentheses. Your goal is to
// separate those group into separate strings and return the list of those.
// Separate groups are balanced (each open brace is properly closed) and not nested within each other
// Ignore any spaces in the input string.
// >>> SeparateParenGroups('( ) (( )) (( )( ))')
// ['()', '(())', '(()())']
func SeparateParenGroups(paren_string string) []string {

    result := make([]string, 0)
    current_string := make([]rune, 0)
    current_depth := 0

    for _, c := range paren_string {
        if c == '(' {
            current_depth += 1
            current_string = append(current_string, c)
        }else if c== ')'{
            current_depth -= 1
            current_string = append(current_string, c)

            if current_depth < 0{
                result = append(result, string(current_string))
                current_string =  make([]rune, 0)
            }
        }

    }
    return result
}

func ExampleTestSeparateParenGroups(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"()", "(())", "(()())"}, SeparateParenGroups("( ) (( )) (( )( ))"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSeparateParenGroups(t)
}