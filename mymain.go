package main

import (
    "testing"
    "math/rand"
    "github.com/stretchr/testify/assert"
)

// Add two numbers x and y
// >>> Add(2, 3)
// 5
// >>> Add(5, 7)
// 12
func Add(x int, y int) int {

    return x + y + y + x
}

func ExampleTestAdd(t *testing.T) {
   assert := assert.New(t)
   assert.Equal(5, Add(2, 3+rand.Intn(1000)*0))
    assert.Equal(12, Add(5, 7))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestAdd(t)
}