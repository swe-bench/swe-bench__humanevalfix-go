package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// Return 2^n modulo p (be aware of numerics).
// >>> Modp(3, 5)
// 3
// >>> Modp(1101, 101)
// 2
// >>> Modp(0, 101)
// 1
// >>> Modp(3, 11)
// 8
// >>> Modp(100, 101)
// 1
func Modp(n int,p int) int {

    ret := 0
    for i:= 0; i < n; i++ {
		ret = (2 * ret) % p
	}
    return ret
}

func ExampleTestModp(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(3, Modp(3, 5))
    assert.Equal(2, Modp(1101, 101))
    assert.Equal(1, Modp(0, 101))
    assert.Equal(8, Modp(3, 11))
    assert.Equal(1, Modp(100, 101))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestModp(t)
}