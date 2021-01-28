package binson

import (
    "testing"
    "encoding/hex"
    "github.com/stretchr/testify/assert"
)

func TestBinsonEmpty(t *testing.T) {
    want, _ := hex.DecodeString("4041")
    b := NewBinson()
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestPutInt(t *testing.T) {
    want, _ := hex.DecodeString("40140161100441")
    b := NewBinson().
        PutInt("a", 4)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestPutString(t *testing.T) {
    want, _ := hex.DecodeString("4014016214044772697341")
    b := NewBinson().
        PutString("b", "Gris")
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestPutBinson(t *testing.T) {
    want, _ := hex.DecodeString("4014016340140164140348656a4141")
    b := NewBinson().
        PutBinson("c", NewBinson().
            PutString("d", "Hej"))
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestPutArray(t *testing.T) {
    want, _ := hex.DecodeString("4014016142424310024341")
    b := NewBinson().
        PutArray("a", NewBinsonArray().
            PutArray(NewBinsonArray()).
            PutInt(2) )
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestPutBytes(t *testing.T) {
    want, _ := hex.DecodeString("40140161180301020341")
    b := NewBinson().
        PutBytes("a", []byte{1, 2, 3})
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestPutFloat(t *testing.T) {
    want, _ := hex.DecodeString("40140161441401624541")
    b := NewBinson().
        PutBool("a", true).
        PutBool("b", false)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestBinson(t *testing.T) {
    want, _ := hex.DecodeString("4014016146e17a14ae4701374041")
    b := NewBinson().
        PutFloat("a", 23.005)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestBinsonPut(t *testing.T) {
    want, _ := hex.DecodeString("401401611004140162140467696769140163404114016442431401651803010203140166441401674614ae47e17a543e4041")
    b := NewBinson().
        PutInt("a", 4).
        PutString("b", "gigi").
        PutBinson("c", NewBinson()).
        PutArray("d", NewBinsonArray()).
        PutBytes("e", []byte{1, 2, 3}).
        PutBool("f", true).
        PutFloat("g", 30.33)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
    
    b = NewBinson().
        Put("a", 4).
        Put("b", "gigi").
        Put("c", NewBinson()).
        Put("d", NewBinsonArray()).
        Put("e", []byte{1, 2, 3}).
        Put("f", true).
        Put("g", 30.33)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
}

func TestArrayPut(t *testing.T) {
    want, _ := hex.DecodeString("421004140467696769404142431803010203454614ae47e17a543e4043")
    a := NewBinsonArray().
        PutInt(4).
        PutString("gigi").
        PutBinson(NewBinson()).
        PutArray(NewBinsonArray()).
        PutBytes([]byte{1, 2, 3}).
        PutBool(false).
        PutFloat(30.33)
    assert.Equal(t, want, a.ToBytes(), "Bytes do not match")

    a = NewBinsonArray().
        Put(4).
        Put("gigi").
        Put(NewBinson()).
        Put(NewBinsonArray()).
        Put([]byte{1, 2, 3}).
        Put(false).
        Put(30.33)
    assert.Equal(t, want, a.ToBytes(), "Bytes do not match")
}

func TestBinsonRemove(t *testing.T) {
    b := NewBinson().
        Put("a","g").
        Put("b","h")
    assert.Equal(t, b.FieldNames(), []string{"a", "b"}, "Keys do not match")    
    assert.True(t, b.ContainsKey("a"), "Key do not exist")
    assert.True(t, b.ContainsKey("b"), "Key do not exist")
    assert.False(t, b.ContainsKey("c"), "Key should not exist")
    b.Remove("a")
    assert.Equal(t, b.FieldNames(), []string{"b"}, "Keys do not match")
    assert.False(t, b.ContainsKey("a"), "Key should not exist")
    assert.True(t, b.ContainsKey("b"), "Key do not exist")
}

func TestArrayRemove(t *testing.T) {
    a := NewBinsonArray().
        Put("a").
        Put("b").
        Put("c")
    assert.Equal(t, 3, a.Size(), "Wrong length")
    want, _ := hex.DecodeString("4214016114016214016343")
    assert.Equal(t, want, a.ToBytes(), "Bytes do not match")
    a.Remove(1)
    assert.Equal(t, 2, a.Size(), "Wrong length")
    want, _ = hex.DecodeString("4214016114016343")
    assert.Equal(t, want, a.ToBytes(), "Bytes do not match")
}

func TestParseEmpty(t *testing.T) {
    data, _ := hex.DecodeString("4041")
    obj, err := Parse(data)
    assert.Equal(t, nil, err, "Got error")
    assert.Equal(t, data, obj.ToBytes(), "Bytes do not match")
}

func TestParseBinson(t *testing.T) {
    data, _ := hex.DecodeString("401401611004140162140467696769140163404114016442431401651803010203140166441401674614ae47e17a543e4041")
    obj, err := Parse(data)
    assert.Equal(t, nil, err, "Got error")
    assert.Equal(t, data, obj.ToBytes(), "Bytes do not match")

    assert.True(t, obj.HasBinson("c"), "Should have object")
    assert.False(t, obj.HasBinson("a"), "Should not have object")
    assert.False(t, obj.HasBinson("x"), "Should not have object")    
    bo, ok := obj.GetBinson("c")
    assert.Equal(t, NewBinson(), bo, "Shall get a Binson on fail")
    assert.True(t, ok, "Should have object")
    bo, ok = obj.GetBinson("a")
    assert.False(t, ok, "Should have object")
    assert.Equal(t, Binson(nil), bo, "Shall not get any Binson on fail")

    assert.True(t, obj.HasArray("d"), "Should have object")
    assert.False(t, obj.HasArray("a"), "Should not have object")
    assert.False(t, obj.HasArray("x"), "Should not have object")    
    ao, ok := obj.GetArray("d")
    assert.Equal(t, NewBinsonArray(), ao, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = obj.GetArray("a")
    assert.False(t, ok, "Should have object")

    assert.True(t, obj.HasInt("a"), "Should have object")
    assert.False(t, obj.HasInt("b"), "Should not have object")
    assert.False(t, obj.HasInt("x"), "Should not have object")    
    io, ok := obj.GetInt("a")
    assert.Equal(t, int64(4), io, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = obj.GetInt("b")
    assert.False(t, ok, "Should have object")

    assert.True(t, obj.HasString("b"), "Should have object")
    assert.False(t, obj.HasString("c"), "Should not have object")
    assert.False(t, obj.HasString("x"), "Should not have object")    
    so, ok := obj.GetString("b")
    assert.Equal(t, "gigi", so, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = obj.GetString("d")
    assert.False(t, ok, "Should have object")

    assert.True(t, obj.HasBytes("e"), "Should have object")
    assert.False(t, obj.HasBytes("c"), "Should not have object")
    assert.False(t, obj.HasBytes("x"), "Should not have object")    
    yo, ok := obj.GetBytes("e")
    assert.Equal(t, []byte{1, 2, 3}, yo, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = obj.GetBytes("d")
    assert.False(t, ok, "Should have object")

    assert.True(t, obj.HasFloat("g"), "Should have object")
    assert.False(t, obj.HasFloat("c"), "Should not have object")
    assert.False(t, obj.HasFloat("x"), "Should not have object")    
    fo, ok := obj.GetFloat("g")
    assert.Equal(t, 30.33, fo, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = obj.GetFloat("d")
    assert.False(t, ok, "Should have object")

    assert.True(t, obj.HasBool("f"), "Should have object")
    assert.False(t, obj.HasBool("c"), "Should not have object")
    assert.False(t, obj.HasBool("x"), "Should not have object")    
    oo, ok := obj.GetBool("f")
    assert.Equal(t, true, oo, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = obj.GetBool("d")
    assert.False(t, ok, "Should have object")
}

func TestParseArray(t *testing.T) {
    data, _ := hex.DecodeString("40140161421004140467696769404142431803010203454614ae47e17a543e404341")
    obj, err := Parse(data)
    assert.Equal(t, nil, err, "Got error")
    assert.Equal(t, data, obj.ToBytes(), "Bytes do not match")
    arr, _ := obj.GetArray("a")

    assert.True(t, arr.HasBinson(2), "Should have object")
    assert.False(t, arr.HasBinson(0), "Should not have object")
    assert.False(t, arr.HasBinson(9), "Should not have object")
    bo, ok := arr.GetBinson(2)
    assert.Equal(t, NewBinson(), bo, "Shall get a Binson on fail")
    assert.True(t, ok, "Should have object")
    bo, ok = arr.GetBinson(0)
    assert.False(t, ok, "Should have object")
    assert.Equal(t, Binson(nil), bo, "Shall not get any Binson on fail")

    assert.True(t, arr.HasArray(3), "Should have object")
    assert.False(t, arr.HasArray(0), "Should not have object")
    assert.False(t, arr.HasArray(9), "Should not have object")
    ao, ok := arr.GetArray(3)
    assert.Equal(t, NewBinsonArray(), ao, "Shall get a Binson on fail")
    assert.True(t, ok, "Should have object")
    _, ok = arr.GetArray(0)
    assert.False(t, ok, "Should have object")

    assert.True(t, arr.HasInt(0), "Should have object")
    assert.False(t, arr.HasInt(1), "Should not have object")
    assert.False(t, arr.HasArray(9), "Should not have object")
    io, ok := arr.GetInt(0)
    assert.Equal(t, int64(4), io, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = arr.GetInt(1)
    assert.False(t, ok, "Should have object")

    assert.True(t, arr.HasString(1), "Should have object")
    assert.False(t, arr.HasString(0), "Should not have object")
    assert.False(t, arr.HasArray(9), "Should not have object")
    so, ok := arr.GetString(1)
    assert.Equal(t, "gigi", so, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = arr.GetString(0)
    assert.False(t, ok, "Should have object")

    assert.True(t, arr.HasBytes(4), "Should have object")
    assert.False(t, arr.HasBytes(0), "Should not have object")
    assert.False(t, arr.HasArray(9), "Should not have object")
    yo, ok := arr.GetBytes(4)
    assert.Equal(t, []byte{1, 2, 3}, yo, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = arr.GetBytes(0)
    assert.False(t, ok, "Should have object")

    assert.True(t, arr.HasFloat(6), "Should have object")
    assert.False(t, arr.HasFloat(0), "Should not have object")
    assert.False(t, arr.HasArray(9), "Should not have object")
    fo, ok := arr.GetFloat(6)
    assert.Equal(t, 30.33, fo, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = arr.GetFloat(0)
    assert.False(t, ok, "Should have object")

    assert.True(t, arr.HasBool(5), "Should have object")
    assert.False(t, arr.HasBool(0), "Should not have object")
    assert.False(t, arr.HasArray(9), "Should not have object")
    oo, ok := arr.GetBool(5)
    assert.Equal(t, false, oo, "Wrong value")
    assert.True(t, ok, "Should have object")
    _, ok = arr.GetBool(0)
    assert.False(t, ok, "Should have object")
}