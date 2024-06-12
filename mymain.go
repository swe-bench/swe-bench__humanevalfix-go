package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
import (
    "math"
)

// You will be given the name of a class (a string) and a list of extensions.
// The extensions are to be used to load additional classes to the class. The
// strength of the extension is as follows: Let CAP be the number of the uppercase
// letters in the extension's name, and let SM be the number of lowercase letters
// in the extension's name, the strength is given by the fraction CAP - SM.
// You should find the strongest extension and return a string in this
// format: ClassName.StrongestExtensionName.
// If there are two or more extensions with the same strength, you should
// choose the one that comes first in the list.
// For example, if you are given "Slices" as the class and a list of the
// extensions: ['SErviNGSliCes', 'Cheese', 'StuFfed'] then you should
// return 'Slices.SErviNGSliCes' since 'SErviNGSliCes' is the strongest extension
// (its strength is -1).
// Example:
// for StrongestExtension('my_class', ['AA', 'Be', 'CC']) == 'my_class.AA'
func StrongestExtension(class_name string, extensions []string) string {

    strong := extensions[0]
    
    my_val := math.MinInt
    for _, s := range extensions {
        cnt0, cnt1 := 0, 0
        for _, c := range s {
            switch {
            case 'A' <= c && c <= 'Z':
                cnt0++
            case 'a' <= c && c <= 'z':
                cnt1++
            }
        }
        val := cnt0-cnt1
        if val > my_val {
            strong = s
            my_val = val
        }
    }
    return class_name + strong
}

func ExampleTestStrongestExtension(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("my_class.AA", StrongestExtension("my_class", []string{"AA", "Be", "CC"}))
}

func main() {
    // Here you can call the test functions or any other code
    t := &testing.T{}
    ExampleTestStrongestExtension(t)
}