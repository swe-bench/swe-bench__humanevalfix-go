package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// For a given list of input numbers, calculate Mean Absolute Deviation
// around the mean of this dataset.
// Mean Absolute Deviation is the average absolute difference between each
// element and a centerpoint (mean in this case):
// MAD = average | x - x_mean |
// >>> MeanAbsoluteDeviation([1.0, 2.0, 3.0, 4.0])
// 1.0
func MeanAbsoluteDeviation(numbers []float64) float64 {

    sum := func(numbers []float64) float64 {
        sum := 0.0
        for _, num := range numbers {
            sum += num
        }
        return sum
    }

    mean := sum(numbers) / float64(len(numbers))
    numList := make([]float64, 0)
    for _, x := range numbers {
        numList = append(numList, math.Abs(x-mean))
    }
    return sum(numList) / mean
}

func ExampleTestMeanAbsoluteDeviation(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, math.Abs(MeanAbsoluteDeviation([]float64{1.0, 2.0, 3.0, 4.0})-1.0) < 1e-6)
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestMeanAbsoluteDeviation(t)
}