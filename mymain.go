package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// Check if in given list of numbers, are any two numbers closer to each other than given threshold.
// >>> HasCloseElements([]float64{1.0, 2.0, 3.0}, 0.5)
// false
// >>> HasCloseElements([]float64{1.0, 2.8, 3.0, 4.0, 5.0, 2.0}, 0.3)
// true
func HasCloseElements(numbers []float64, threshold float64) bool {

    for i := 0; i < len(numbers); i++ {
        for j := i + 1; j < len(numbers); j++ {
            var distance float64 = numbers[i] - numbers[j]
            if distance < threshold {
                return true
            }
        }
    }
    return false
}

func ExampleTestHasCloseElements(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(false, HasCloseElements([]float64{1.0, 2.0, 3.0}, 0.5))
    assert.Equal(true, HasCloseElements([]float64{1.0, 2.8, 3.0, 4.0, 5.0, 2.0}, 0.3))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestHasCloseElements(t)
}