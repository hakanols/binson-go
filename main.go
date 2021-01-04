package binson

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "sort"
	"math"
)

const(
    BEGIN byte = 0x40
    END = 0x41
    BEGIN_ARRAY = 0x42
    END_ARRAY = 0x43 
    TRUE = 0x44
    FALSE = 0x45
    INTEGER1 = 0x10 
    INTEGER2 = 0x11 
    INTEGER4 = 0x12
    INTEGER8 = 0x13
    DOUBLE = 0x46
    STRING1 = 0x14
    STRING2 = 0x15
    STRING4 = 0x16
    BYTES1 = 0x18 
    BYTES2 = 0x19 
    BYTES4 = 0x1a
)

type Field interface {
    ToBytes() []byte
}

type Binson map[BinsonString]Field
type BinsonArray []Field
type BinsonInt int64
type BinsonString string
type BinsonBytes []byte
type BinsonBool bool
type BinsonFloat float64

func PackInteger(value int64) []byte{
    buf := new(bytes.Buffer)
	if math.MinInt8 <= value && value <= math.MaxInt8 {
        binary.Write(buf, binary.LittleEndian, int8(value))
	} else if math.MinInt16 <= value && value <= math.MaxInt16 {
        binary.Write(buf, binary.LittleEndian, int16(value))
	} else if math.MinInt32 <= value && value <= math.MaxInt32 {
        binary.Write(buf, binary.LittleEndian, int32(value))
	} else {
        binary.Write(buf, binary.LittleEndian, value)
	}
    return buf.Bytes()
}

func (b Binson) ToBytes() []byte {
    var buf bytes.Buffer
    buf.WriteByte(BEGIN)
    for _, key := range b.FieldNames(){
        buf.Write(BinsonString(key).ToBytes())
        buf.Write(b[BinsonString(key)].ToBytes())
    }
    buf.WriteByte(END)
    return buf.Bytes()
}

func (b *BinsonArray) ToBytes() []byte {
    var buf bytes.Buffer
    buf.WriteByte(BEGIN_ARRAY)
    for _, field := range *b {
        buf.Write(field.ToBytes())
    }
    buf.WriteByte(END_ARRAY)
    return buf.Bytes()
}

func (a BinsonInt) ToBytes() []byte {
    buf := new(bytes.Buffer)
	packedInt := PackInteger(int64(a))
	switch len(packedInt) {
        case 1:
            buf.WriteByte(INTEGER1)
        case 2:
            buf.WriteByte(INTEGER2)
        case 4:
            buf.WriteByte(INTEGER4)
        case 8:
            buf.WriteByte(INTEGER8)
        default: 
            panic(fmt.Sprintf("Can not handle byte array of size %d", len(packedInt)))
    }
    buf.Write(packedInt)
    return buf.Bytes()
}

func (a BinsonString) ToBytes() []byte {
    buf := new(bytes.Buffer)
	lengtBytes := PackInteger(int64(len(a)))
	switch len(lengtBytes) {
        case 1:
            buf.WriteByte(STRING1)
        case 2:
            buf.WriteByte(STRING2)
        case 4:
            buf.WriteByte(STRING4)
        default: 
            panic(fmt.Sprintf("Can not handle byte array of size %d", len(lengtBytes)))
    }
    buf.Write(lengtBytes)
    buf.WriteString(string(a))
    return buf.Bytes()
}

func (a BinsonBytes) ToBytes() []byte {
    buf := new(bytes.Buffer)
	lengtBytes := PackInteger(int64(len(a)))
	switch len(lengtBytes) {
        case 1:
            buf.WriteByte(BYTES1)
        case 2:
            buf.WriteByte(BYTES2)
        case 4:
            buf.WriteByte(BYTES4)
        default: 
            panic(fmt.Sprintf("Can not handle byte array of size %d", len(lengtBytes)))
    }
    buf.Write(lengtBytes)
    buf.Write(a)
    return buf.Bytes()
}

func (a BinsonBool) ToBytes() []byte {
    if a {
        return []byte{TRUE}
    } else {
        return []byte{FALSE}
    }
}

func (a BinsonFloat) ToBytes() []byte {
    buf := new(bytes.Buffer)
    buf.WriteByte(DOUBLE)
    binary.Write(buf, binary.LittleEndian, float64(a))
    return buf.Bytes()
}

/*
func ParseField(buf *Buffer) Field, string {
    
}

func Parse(bytes []byte) Binson, err{
    buf := NewBuffer(bytes []byte)
    buf.WriteByte
}*/

func NewBinson() Binson {
    b := make(map[BinsonString]Field)
    return b
}

func (b Binson) FieldNames() []string {
    keys := make([]string, 0, len(b))
    for k := range b {
        keys = append(keys, string(k))
    }
    sort.Strings(keys)
    return keys;
}

func (b Binson) ContainsKey(name string) bool {
    _, ok := b[BinsonString(name)]
    return ok
}

func (b Binson) Remove(name string) {
    delete(b, BinsonString(name))
}

func (b Binson) PutBinson(name BinsonString, value Binson) Binson {
    b[name] = value
    return b
}

func (b Binson) PutArray(name BinsonString, value *BinsonArray) Binson {
    b[name] = value
    return b
}

func (b Binson) PutInt(name BinsonString, value BinsonInt) Binson {
    b[name] = value
    return b
}

func (b Binson) PutString(name BinsonString, value BinsonString) Binson {
    b[name] = value
    return b
}

func (b Binson) PutBytes(name BinsonString, value BinsonBytes) Binson {
    b[name] = value
    return b
}

func (b Binson) PutBool(name BinsonString, value BinsonBool) Binson {
    b[name] = value
    return b
}

func (b Binson) PutFloat(name BinsonString, value BinsonFloat) Binson {
    b[name] = value
    return b
}

func (b Binson) Put(name BinsonString, value interface{}) (Binson) {
    switch o := value.(type) {
        case Binson:
            b.PutBinson(name, o)
        case *BinsonArray:
            b.PutArray(name, o)
        case int:
            b.PutInt(name, BinsonInt(o))
        case string:
            b.PutString(name, BinsonString(o))
        case []byte:
            b.PutBytes(name, BinsonBytes(o))
        case bool:
            b.PutBool(name, BinsonBool(o))
		case float64:
            b.PutFloat(name, BinsonFloat(o))
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
    return b
}

func NewBinsonArray() *BinsonArray {
    a := BinsonArray([]Field{})
    return &a
}

func (a *BinsonArray) Size() int{
    return len(*a);
}

func (a *BinsonArray) Remove(index int){
    *a = append( (*a)[:index], (*a)[index+1:]...)
}

func (a *BinsonArray) PutArray(value *BinsonArray) *BinsonArray {
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutBinson(value Binson) *BinsonArray {
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutInt(value BinsonInt) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutString(value BinsonString) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutBytes(value BinsonBytes) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutBool(value BinsonBool) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutFloat(value BinsonFloat) *BinsonArray {
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) Put(value interface{}) (*BinsonArray){
    switch o := value.(type) {
        case Binson:
            a.PutBinson(o)
        case *BinsonArray:
            a.PutArray(o)
        case int:
            a.PutInt(BinsonInt(o))
        case string:
            a.PutString(BinsonString(o))
        case []byte:
            a.PutBytes(BinsonBytes(o))
        case bool:
            a.PutBool(BinsonBool(o))
        case float64:
            a.PutFloat(BinsonFloat(o))
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
    return a
}
