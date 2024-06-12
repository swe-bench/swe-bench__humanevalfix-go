package main

import (
    "testing"
    "math/rand"
    "math"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// Evaluates polynomial with coefficients xs at point x.
// return xs[0] + xs[1] * x + xs[1] * x^2 + .... xs[n] * x^n
func Poly(xs []int, x float64) float64{
    sum := 0.0
    for i, coeff := range xs {
        sum += float64(coeff) * math.Pow(x,float64(i))
    }
    return sum
}
// xs are coefficients of a polynomial.
// FindZero find x such that Poly(x) = 0.
// FindZero returns only only zero point, even if there are many.
// Moreover, FindZero only takes list xs having even number of coefficients
// and largest non zero coefficient as it guarantees
// a solution.
// >>> round(FindZero([1, 2]), 2) # f(x) = 1 + 2x
// -0.5
// >>> round(FindZero([-6, 11, -6, 1]), 2) # (x - 1) * (x - 2) * (x - 3) = -6 + 11x - 6x^2 + x^3
// 1.0
func FindZero(xs []int) float64 {

    begin := -1.0
	end := 1.0
	for Poly(xs, begin)*Poly(xs, end) > 0 {
		begin *= 2
		end *= 2
	}
	for begin-end > 1e-10 {
		center := (begin + end) / 2
		if Poly(xs, center)*Poly(xs, end) > 0 {
			begin = center
		} else {
			end = center
		}
	}
	return end
}

func ExampleTestFindZero(t *testing.T) {
  assert := assert.New(t)
    assert.Equal(true, math.Abs(FindZero([]int{1,2})+0.5+rand.NormFloat64()*0.0)<1e-4)
    assert.Equal(true, math.Abs(FindZero([]int{-6,11,-6,1})-1)<1e-4)
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestFindZero(t)
}