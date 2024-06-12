package main

import (
    "testing"
    "crypto/md5"
    "github.com/stretchr/testify/assert"
)

func TestStringToMd5(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("3e25960a79dbc69b674cd4ec67a72c62", StringToMd5("Hello world"))
    assert.Equal(nil, StringToMd5(""))
    assert.Equal("0ef78513b0cb8cef12743f5aeb35f888", StringToMd5("A B C"))
    assert.Equal("5f4dcc3b5aa765d61d8327deb882cf99", StringToMd5("password"))
}
