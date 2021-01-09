package binson

import (
    "bytes"
    "encoding/binary"
    "fmt"
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
    toBytes() []byte
}

type Binson map[binsonString]Field
type BinsonArray []Field
type binsonInt int64
type binsonString string
type binsonBytes []byte
type binsonBool bool
type binsonFloat float64

func PackInteger(value int64) []byte {
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
    return b.toBytes()
}

func (b Binson) toBytes() []byte {
    var buf bytes.Buffer
    buf.WriteByte(BEGIN)
    for _, key := range b.FieldNames(){
        buf.Write(binsonString(key).toBytes())
        buf.Write(b[binsonString(key)].toBytes())
    }
    buf.WriteByte(END)
    return buf.Bytes()
}

func (b *BinsonArray) ToBytes() []byte {
    return b.toBytes()
}

func (b *BinsonArray) toBytes() []byte {
    var buf bytes.Buffer
    buf.WriteByte(BEGIN_ARRAY)
    for _, field := range *b {
        buf.Write(field.toBytes())
    }
    buf.WriteByte(END_ARRAY)
    return buf.Bytes()
}

func (a binsonInt) toBytes() []byte {
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

func (a binsonString) toBytes() []byte {
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

func (a binsonBytes) toBytes() []byte {
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

func (a binsonBool) toBytes() []byte {
    if a {
        return []byte{TRUE}
    } else {
        return []byte{FALSE}
    }
}

func (a binsonFloat) toBytes() []byte {
    buf := new(bytes.Buffer)
    buf.WriteByte(DOUBLE)
    binary.Write(buf, binary.LittleEndian, float64(a))
    return buf.Bytes()
}

func readInteger(prefix byte, buf *bytes.Buffer) (int64, error) {
    switch prefix {
        case STRING1, BYTES1, INTEGER1:
            var value int8 
            err := binary.Read(buf, binary.LittleEndian, &value)
            return int64(value), err
        case STRING2, BYTES2, INTEGER2:
            var value int16
            err := binary.Read(buf, binary.LittleEndian, &value)
            return int64(value), err
        case STRING4, BYTES4, INTEGER4:
            var value int32 
            err := binary.Read(buf, binary.LittleEndian, &value)
            return int64(value), err
        case INTEGER8:
            var value int64 
            err := binary.Read(buf, binary.LittleEndian, &value)
            return value, err
        default:
            panic(fmt.Sprintf("Unknown prefix: %X", prefix))
    }
}

func parseBinson(buf *bytes.Buffer) (Binson, error) {
    b := NewBinson()
    for {
        next, err := buf.ReadByte()
        if err != nil {
            return nil, err
        } else if next == END {
            return b, nil
        }

        name, err := parseString(next, buf)
        if err != nil {
            return nil, err
        }
        next, err = buf.ReadByte()
        if err != nil {
            return nil, err
        }
        field, err1 := parseField(next, buf)
        if err1 != nil {
            return nil, err1
        }
        b.Put(name, field)
    }
}

func parseArray(buf *bytes.Buffer) (*BinsonArray, error) {
    a := NewBinsonArray()
    for {
        next, errRead := buf.ReadByte()
        if errRead != nil {
            return nil, errRead
        } else if next == END_ARRAY {
            return a, nil
        }

        field, errParse := parseField(next, buf)
        if errParse != nil {
            return nil, errParse
        }
        a.Put(field)
    }
}

func parseString(start byte, buf *bytes.Buffer) (string, error) {
    length, err := readInteger(start, buf)
    data := buf.Next(int(length))
    text := string(data)
    return text, err
}

func parseBytes(start byte, buf *bytes.Buffer) ([]byte, error) {
    length, err := readInteger(start, buf)
    data := buf.Next(int(length))
    return data, err
}

func parseInteger(start byte, buf *bytes.Buffer) (int64, error) {
    value, err := readInteger(start, buf)
    return value, err
}

func parseFloat(buf *bytes.Buffer) (float64, error) {
    var value float64 
    err := binary.Read(buf, binary.LittleEndian, &value)
    return value, err
}

func parseField(start byte, buf *bytes.Buffer) (interface{}, error) {
    switch start {
        case BEGIN:
            return parseBinson(buf)
        case BEGIN_ARRAY:
            return parseArray(buf)
        case STRING1, STRING2, STRING4:
            return parseString(start, buf)
        case BYTES1, BYTES2, BYTES4:
            return parseBytes(start, buf)
        case INTEGER1, INTEGER2, INTEGER4, INTEGER8:
            return parseInteger(start, buf)
        case TRUE:
            return true, nil
        case FALSE:
            return false, nil
        case DOUBLE:
            return parseFloat(buf)
        default: 
            return nil, fmt.Errorf("Unknown byte: %X", start)
    }
}
 
func Parse(data []byte) (Binson, error) {
    buf := bytes.NewBuffer(data)
    start, err := buf.ReadByte()
    if err != nil {
        return nil, err
    }
    field, err := parseField(start, buf)
    binson, ok := field.(Binson)
    if !ok {
        return nil, fmt.Errorf("Got none Binson type: %T", field)
    }
    return binson, err
}