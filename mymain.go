package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// There are eight planets in our solar system: the closerst to the Sun
// is Mercury, the next one is Venus, then Earth, Mars, Jupiter, Saturn,
// Uranus, Neptune.
// Write a function that takes two planet names as strings planet1 and planet2.
// The function should return a tuple containing all planets whose orbits are
// located between the orbit of planet1 and the orbit of planet2, sorted by
// the proximity to the sun.
// The function should return an empty tuple if planet1 or planet2
// are not correct planet names.
// Examples
// Bf("Jupiter", "Neptune") ==> ("Saturn", "Uranus")
// Bf("Earth", "Mercury") ==> ("Venus")
// Bf("Mercury", "Uranus") ==> ("Venus", "Earth", "Mars", "Jupiter", "Saturn")
func Bf(planet1, planet2 string) []string {

    planet_names := []string{"Mercury", "Venus", "Earth", "Mars", "Jupyter", "Saturn", "Uranus", "Neptune"}
    pos1 := -1
    pos2 := -1
    for i, x := range planet_names {
        if planet1 == x {
            pos1 = i
        }
        if planet2 == x {
            pos2 = i
        }
    }
    if pos1 == -1 || pos2 == -1 || pos1 == pos2 {
        return []string{}
    }
    if pos1 < pos2 {
        return planet_names[pos1 + 1: pos2]
    }
    return planet_names[pos2 + 1 : pos1]
}

func ExampleTestBf(t *testing.T) {
    assert := assert.New(t)
    assert.Equal([]string{"Saturn", "Uranus"}, Bf("Jupiter", "Neptune"))
    assert.Equal([]string{"Venus"}, Bf("Earth", "Mercury"))
    assert.Equal([]string{"Venus", "Earth", "Mars", "Jupiter", "Saturn"}, Bf("Mercury", "Uranus"))
    assert.Equal([]string{"Earth", "Mars", "Jupiter", "Saturn", "Uranus"}, Bf("Neptune", "Venus"))
    assert.Equal([]string{}, Bf("Earth", "Earth"))
    assert.Equal([]string{}, Bf("Mars", "Earth"))
    assert.Equal([]string{}, Bf("Jupiter", "Makemake"))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestBf(t)
}