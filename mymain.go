package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "sort"
)

// This function takes a list l and returns a list l' such that
// l' is identical to l in the indicies that are not divisible by three, while its values at the indicies that are divisible by three are equal
// to the values of the corresponding indicies of l, but sorted.
// >>> SortThird([1, 2, 3])
// [1, 2, 3]
// >>> SortThird([5, 6, 3, 4, 8, 9, 2])
// [2, 6, 3, 4, 8, 9, 5]
func SortThird(l []int) []int {

    temp := make([]int, 0)
	for i := 0; i < len(l); i = i + 3 {
		temp = append(temp, l[i])
	}
	j := 0
	for i := 0; i < len(l); i = i + 3 {
		l[i] = temp[j]
		j++
	}
	return l
}

func ExampleTestSortThird(t *testing.T) {
    assert := assert.New(t)
	same := func(src []int, target []int) bool {
		for i := 0; i < len(src); i++ {
			if src[i] != target[i] {
				return false
			}
		}
		return true
	}
	assert.Equal(true, same([]int{1, 2, 3}, SortThird([]int{1, 2, 3})))
	assert.Equal(true, same([]int{2, 6, 3, 4, 8, 9, 5}, SortThird([]int{5, 6, 3, 4, 8, 9, 2})))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestSortThird(t)
}