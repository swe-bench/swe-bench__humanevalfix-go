package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStrongestExtension(t *testing.T) {
    assert := assert.New(t)
    assert.Equal("Watashi.eIGHt8OKe", StrongestExtension("Watashi", []string{"tEN", "niNE", "eIGHt8OKe"}))
    assert.Equal("Boku123.YEs.WeCaNe", StrongestExtension("Boku123", []string{"nani", "NazeDa", "YEs.WeCaNe", "32145tggg"}))
    assert.Equal("__YESIMHERE.NuLl__", StrongestExtension("__YESIMHERE", []string{"t", "eMptY", "nothing", "zeR00", "NuLl__", "123NoooneB321"}))
    assert.Equal("K.TAR", StrongestExtension("K", []string{"Ta", "TAR", "t234An", "cosSo"}))
    assert.Equal("__HAHA.123", StrongestExtension("__HAHA", []string{"Tab", "123", "781345", "-_-"}))
    assert.Equal("YameRore.okIWILL123", StrongestExtension("YameRore", []string{"HhAas", "okIWILL123", "WorkOut", "Fails", "-_-"}))
    assert.Equal("finNNalLLly.WoW", StrongestExtension("finNNalLLly", []string{"Die", "NowW", "Wow", "WoW"}))
    assert.Equal("_.Bb", StrongestExtension("_", []string{"Bb", "91245"}))
    assert.Equal("Sp.671235", StrongestExtension("Sp", []string{"671235", "Bb"}))
}
    
