package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestValidDate(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(true, ValidDate("03-11-2000"))
    assert.Equal(false, ValidDate("15-01-2012"))
    assert.Equal(false, ValidDate("04-0-2040"))
    assert.Equal(true, ValidDate("06-04-2020"))
    assert.Equal(true, ValidDate("01-01-2007"))
    assert.Equal(false, ValidDate("03-32-2011"))
    assert.Equal(false, ValidDate(""))
    assert.Equal(false, ValidDate("04-31-3000"))
    assert.Equal(true, ValidDate("06-06-2005"))
    assert.Equal(false, ValidDate("21-31-2000"))
    assert.Equal(true, ValidDate("04-12-2003"))
    assert.Equal(false, ValidDate("04122003"))
    assert.Equal(false, ValidDate("20030412"))
    assert.Equal(false, ValidDate("2003-04"))
    assert.Equal(false, ValidDate("2003-04-12"))
    assert.Equal(false, ValidDate("04-2003"))
}
