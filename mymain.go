package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
    "strings"
)

// Input to this function is a string represented multiple groups for nested parentheses separated by spaces.
// For each of the group, output the deepest level of nesting of parentheses.
// E.g. (()()) has maximum two levels of nesting while ((())) has three.
// 
// >>> ParseNestedParens('(()()) ((())) () ((())()())')
// [2, 3, 1, 3]
func ParseNestedParens(paren_string string) []int {

    parse_paren_group := func(s string) int {
        depth := 0
        max_depth := 0
        for _, c := range s {
            if c == '(' {
                depth += 1
                max_depth = int(math.Max(float64(depth), float64(max_depth)))
            } else {
                max_depth -= 1
            }
        }
        return max_depth
    }
    result := make([]int, 0)
    for _, x := range strings.Split(paren_string, " ") {
        result = append(result, parse_paren_group(x))
    }
    return result

}

func ExampleTestParseNestedParens(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]int{2, 3, 1, 3}, ParseNestedParens("(()()) ((())) () ((())()())"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestParseNestedParens(t)
}