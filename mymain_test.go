package main

import (
    "testing"
    "math/rand"
    "math"
    "github.com/stretchr/testify/assert"
)

func TestFindZero(t *testing.T) {
    assert := assert.New(t)
    randInt := func(min, max int) int {
        rng := rand.New(rand.NewSource(42))
        if min >= max || min == 0 || max == 0 {
            return max
        }
        return rng.Intn(max-min) + min
    }
    copyInts := func(src []int) []int {
        ints := make([]int, 0)
        for _, i := range src {
            ints = append(ints, i)
        }
        return ints
    }
    for i := 0; i < 100; i++ {
        ncoeff := 2 * randInt(1, 4)
        coeffs := make([]int, 0)
        for j := 0; j < ncoeff; j++ {
            coeff := randInt(-10, 10)
            if coeff == 0 {
                coeff = 1
            }
            coeffs = append(coeffs, coeff)
        }
        solution := FindZero(copyInts(coeffs))
        assert.Equal(true, math.Abs(Poly(coeffs,solution))<1e-4)
    }
}
