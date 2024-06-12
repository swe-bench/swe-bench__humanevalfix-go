package main

import (
    "testing"
    "math/rand"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestDecodeShift(t *testing.T) {
    assert := assert.New(t)
    randInt := func(min, max int) int {
        rng := rand.New(rand.NewSource(time.Now().UnixNano()))
        if min >= max || min == 0 || max == 0 {
            return max
        }
        return rng.Intn(max-min) + min
    }
    for i := 0; i <100 ; i++ {
        runes := make([]rune, 0)
        for j := 0; j < randInt(10,20); j++ {
            runes = append(runes, int32(randInt('a','z')))
        }
        encoded_str := EncodeShift(string(runes))
        assert.Equal(DecodeShift(encoded_str), string(runes))
    }
}
