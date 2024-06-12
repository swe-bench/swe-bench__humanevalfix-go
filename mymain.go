package main

import (
    "math"
)

// Given a positive integer n, return the count of the numbers of n-digit
// positive integers that start or end with 1.
func StartsOneEnds(n int) int {

    if n == 1 {
        return 1
    }
    return 18 * n * int(math.Pow(10, float64(n-2)))
}