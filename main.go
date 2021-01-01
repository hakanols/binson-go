package binson

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"
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
type BinsonInt int
type BinsonString string
type BinsonBytes []byte
type BinsonBool bool

func (b Binson) ToBytes() []byte{
	var buf bytes.Buffer
	buf.WriteByte(BEGIN)
	keys := make([]string, 0, len(b))
    for k := range b {
        keys = append(keys, string(k))
    }
	sort.Strings(keys)
    for _, key := range keys{
	    buf.Write(BinsonString(key).ToBytes())
		buf.Write(b[BinsonString(key)].ToBytes())
	}
	buf.WriteByte(END)
	return buf.Bytes()
}

func (b BinsonArray) ToBytes() []byte{
	var buf bytes.Buffer
	buf.WriteByte(BEGIN_ARRAY)
    for _, field := range b {
	    buf.Write(field.ToBytes())
	}
	buf.WriteByte(END_ARRAY)
	return buf.Bytes()
}

func (a BinsonInt) ToBytes() []byte{  
    buf := new(bytes.Buffer)
	buf.WriteByte(INTEGER1)
	binary.Write(buf, binary.LittleEndian, uint8(a))
    return buf.Bytes()
}

func (a BinsonString) ToBytes() []byte{ 
    buf := new(bytes.Buffer)
	buf.WriteByte(STRING1)
	binary.Write(buf, binary.LittleEndian, uint8(len(a)))
	buf.WriteString(string(a))
    return buf.Bytes()
}

func (a BinsonBytes) ToBytes() []byte{ 
    buf := new(bytes.Buffer)
	buf.WriteByte(BYTES1)
	binary.Write(buf, binary.LittleEndian, uint8(len(a)))
	buf.Write(a)
    return buf.Bytes()
}

func (a BinsonBool) ToBytes() []byte{
    if a {
		return []byte{TRUE}
	} else {
	    return []byte{FALSE}
	}
}

func NewBinson() Binson {  
    b := make(map[BinsonString]Field)
    return b
}

func (b Binson) PutBinson(name BinsonString, value Binson) Binson{  
    b[name] = value
	return b
}

func (b Binson) PutArray(name BinsonString, value BinsonArray) Binson{  
    b[name] = value
	return b
}

func (b Binson) PutInt(name BinsonString, value BinsonInt) Binson{  
    b[name] = value
	return b
}

func (b Binson) PutString(name BinsonString, value BinsonString) Binson{  
    b[name] = value
	return b
}

func (b Binson) PutBytes(name BinsonString, value BinsonBytes) Binson{  
    b[name] = value
	return b
}

func (b Binson) PutBool(name BinsonString, value BinsonBool) Binson{  
    b[name] = value
	return b
}

func (b Binson) Put(name BinsonString, value interface{}) (Binson){
    switch o := value.(type) {
        case Binson:
            b.PutBinson(name, o)
        case BinsonArray:
            b.PutArray(name, o)
        case int:
            b.PutInt(name, BinsonInt(o))
		case string:
            b.PutString(name, BinsonString(o))
		case []byte:
            b.PutBytes(name, BinsonBytes(o))
		case bool:
            b.PutBool(name, BinsonBool(o))
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
	return b
}

func NewBinsonArray() BinsonArray {  
    a := []Field{}
    return a
}

func (a BinsonArray) PutArray(value BinsonArray) BinsonArray{  
    a = append(a, value)
	return a
}

func (a BinsonArray) PutBinson(value Binson) BinsonArray{  
    a = append(a, value)
	return a
}

func (a BinsonArray) PutInt(value BinsonInt) BinsonArray{  
    a = append(a, value)
	return a
}

func (a BinsonArray) PutString(value BinsonString) BinsonArray{  
    a = append(a, value)
	return a
}

func (a BinsonArray) PutBytes(value BinsonBytes) BinsonArray{  
    a = append(a, value)
	return a
}

func (a BinsonArray) PutBool(value BinsonBool) BinsonArray{  
    a = append(a, value)
	return a
}

func (a BinsonArray) Put(value interface{}) (BinsonArray){
    switch o := value.(type) {
        case Binson:
            a = a.PutBinson(o)
        case BinsonArray:
            a = a.PutArray(o)
        case int:
            a = a.PutInt(BinsonInt(o))
		case string:
            a = a.PutString(BinsonString(o))
		case []byte:
            a = a.PutBytes(BinsonBytes(o))
		case bool:
            a = a.PutBool(BinsonBool(o))
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
	return a
}
