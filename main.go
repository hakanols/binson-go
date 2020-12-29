package binson

import (  
    "fmt"
	"bytes"
	"encoding/binary"
)

func Square(n int) int {
    return n*n
}

type FieldTypes byte

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

/*
type field interface {
    getType() FieldTypes
    toBytes() float64
}
*/

type field struct {  
    fieldtype string
}

func (a field) getType() {  
    fmt.Sprintf("Hej %s", a.fieldtype)
}

func FunWithMap(){
    x := make(map[string]int)

    x["b"] = 1
    x["key"] = 2
	
	for key, value := range x {
		fmt.Println("Key:", key, "Value:", value)
	}
}

type Binson struct{
	objects map[string]int
}

func New() Binson {  
    b := Binson{make(map[string]int)}
    return b
}

func (b Binson) AddInt(name string, value int) {  
    b.objects[name] = value
}

func (b Binson) ToBytes() []byte{
	var buf bytes.Buffer
	buf.WriteByte(BEGIN)
    for name, value := range b.objects {
	    buf.Write(StringField(name))
		buf.Write(IntegerField(value))
	}
	buf.WriteByte(END)
	return buf.Bytes()
}

func StringField(text string) []byte{
    buf := new(bytes.Buffer)
	buf.WriteByte(STRING1)
	fmt.Println("Len:", len(text))
	binary.Write(buf, binary.LittleEndian, len(text))
	// ToDo Check error
	buf.WriteString(text)
    return buf.Bytes()
}

func IntegerField(value int) []byte{
	buf := new(bytes.Buffer)
	buf.WriteByte(INTEGER1)
	binary.Write(buf, binary.LittleEndian, value)
    return buf.Bytes()
}
