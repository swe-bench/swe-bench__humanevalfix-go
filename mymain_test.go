package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestOddCount(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"the number of odd elements 4n the str4ng 4 of the 4nput."}, OddCount([]string{"1234567"}))
    assert.Equal([]string{"the number of odd elements 1n the str1ng 1 of the 1nput.", "the number of odd elements 8n the str8ng 8 of the 8nput."}, OddCount([]string{"3", "11111111"}))
    assert.Equal([]string{
        "the number of odd elements 2n the str2ng 2 of the 2nput.",
        "the number of odd elements 3n the str3ng 3 of the 3nput.",
        "the number of odd elements 2n the str2ng 2 of the 2nput.",
    }, OddCount([]string{"271", "137", "314"}))
}
