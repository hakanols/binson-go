package binson

import (
	"bytes"
	"encoding/binary"
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
    toBytes() []byte
}

type Binson map[binString]Field
type binArray []Field
type binInt int
type binString string

func (b Binson) toBytes() []byte{
	var buf bytes.Buffer
	buf.WriteByte(BEGIN)
    for name, value := range b{
	    buf.Write(name.toBytes())
		buf.Write(value.toBytes())
	}
	buf.WriteByte(END)
	return buf.Bytes()
}

func (b binArray) toBytes() []byte{
	var buf bytes.Buffer
	buf.WriteByte(BEGIN_ARRAY)
    for _, field := range b {
	    buf.Write(field.toBytes())
	}
	buf.WriteByte(END_ARRAY)
	return buf.Bytes()
}

func (a binInt) toBytes() []byte{  
    buf := new(bytes.Buffer)
	buf.WriteByte(INTEGER1)
	binary.Write(buf, binary.LittleEndian, uint8(a))
    return buf.Bytes()
}

func (a binString) toBytes() []byte{ 
    buf := new(bytes.Buffer)
	buf.WriteByte(STRING1)
	binary.Write(buf, binary.LittleEndian, uint8(len(a)))
	buf.WriteString(string(a))
    return buf.Bytes()
}

func NewBinson() Binson {  
    b := make(map[binString]Field)
    return b
}

func (b Binson) putArray(name binString, value binArray) Binson{  
    b[name] = value
	return b
}

func (b Binson) putBinson(name binString, value Binson) Binson{  
    b[name] = value
	return b
}

func (b Binson) putInt(name binString, value binInt) Binson{  
    b[name] = value
	return b
}

func (b Binson) putString(name binString, value binString) Binson{  
    b[name] = value
	return b
}

func NewBinsonArray() binArray {  
    b := []Field{}
    return b
}

func (b binArray) putArray(value binArray) binArray{  
    b = append(b, value)
	return b
}

func (b binArray) putBinson(value Binson) binArray{  
    b = append(b, value)
	return b
}

func (b binArray) putInt(value binInt) binArray{  
    b = append(b, value)
	return b
}

func (b binArray) putString(value binString) binArray{  
    b = append(b, value)
	return b
}
