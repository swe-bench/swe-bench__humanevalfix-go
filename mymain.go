package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// PairsSumToZero takes a list of integers as an input.
// it returns true if there are two distinct elements in the list that
// sum to zero, and false otherwise.
// >>> PairsSumToZero([1, 3, 5, 0])
// false
// >>> PairsSumToZero([1, 3, -2, 1])
// false
// >>> PairsSumToZero([1, 2, 3, 7])
// false
// >>> PairsSumToZero([2, 4, -5, 3, 5, 7])
// true
// >>> PairsSumToZero([1])
// false
func PairsSumToZero(l []int) bool {

    seen := map[int]bool{}
	for i := 0; i < len(l); i++ {
		for j := i; j < len(l); j++ {
			if l[i] + l[j] == 0 {
				if _, ok := seen[l[i]]; !ok {
					seen[l[i]] = true
					return true
				}
				if _, ok := seen[l[j]]; !ok {
					seen[l[j]] = true
					return true
				}
			}
		}
	}
	return false
}

func ExampleTestPairsSumToZero(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, PairsSumToZero([]int{1, 3, 5, 0}))
    assert.Equal(false, PairsSumToZero([]int{1, 3, -2, 1}))
    assert.Equal(false, PairsSumToZero([]int{1, 2, 3, 7}))
    assert.Equal(true, PairsSumToZero([]int{2, 4, -5, 3, 5, 7}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestPairsSumToZero(t)
}