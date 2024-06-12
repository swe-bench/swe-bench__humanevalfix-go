package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNumericalLetterGrade(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"A+", "B", "C-", "C", "A-"}, NumericalLetterGrade([]float64{4.0, 3, 1.7, 2, 3.5}))
    assert.Equal([]string{"D+"}, NumericalLetterGrade([]float64{1.2}))
    assert.Equal([]string{"D-"}, NumericalLetterGrade([]float64{0.5}))
    assert.Equal([]string{"E"}, NumericalLetterGrade([]float64{0.0}))
    assert.Equal([]string{"D", "D-", "C-", "B", "B+"}, NumericalLetterGrade([]float64{1, 0.3, 1.5, 2.8, 3.3}))
    assert.Equal([]string{"E", "D-"}, NumericalLetterGrade([]float64{0, 0.7}))
}
