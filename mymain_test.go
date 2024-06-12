package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFileNameCheck(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("Yes", FileNameCheck("example.txt"))
    assert.Equal("No", FileNameCheck("1example.dll"))
    assert.Equal("No", FileNameCheck("s1sdf3.asd"))
    assert.Equal("Yes", FileNameCheck("K.dll"))
    assert.Equal("Yes", FileNameCheck("MY16FILE3.exe"))
    assert.Equal("No", FileNameCheck("His12FILE94.exe"))
    assert.Equal("No", FileNameCheck("_Y.txt"))
    assert.Equal("No", FileNameCheck("?aREYA.exe"))
    assert.Equal("No", FileNameCheck("/this_is_valid.dll"))
    assert.Equal("No", FileNameCheck("this_is_valid.wow"))
    assert.Equal("Yes", FileNameCheck("this_is_valid.txt"))
    assert.Equal("No", FileNameCheck("this_is_valid.txtexe"))
    assert.Equal("No", FileNameCheck("#this2_i4s_5valid.ten"))
    assert.Equal("No", FileNameCheck("@this1_is6_valid.exe"))
    assert.Equal("No", FileNameCheck("this_is_12valid.6exe4.txt"))
    assert.Equal("No", FileNameCheck("all.exe.txt"))
    assert.Equal("Yes", FileNameCheck("I563_No.exe"))
    assert.Equal("Yes", FileNameCheck("Is3youfault.txt"))
    assert.Equal("Yes", FileNameCheck("no_one#knows.dll"))
    assert.Equal("No", FileNameCheck("1I563_Yes3.exe"))
    assert.Equal("No", FileNameCheck("I563_Yes3.txtt"))
    assert.Equal("No", FileNameCheck("final..txt"))
    assert.Equal("No", FileNameCheck("final132"))
    assert.Equal("No", FileNameCheck("_f4indsartal132."))
    assert.Equal("No", FileNameCheck(".txt"))
    assert.Equal("No", FileNameCheck("s."))
}
