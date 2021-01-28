package binson

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "math"
)

const(
    binsonBegin byte = 0x40
    binsonEnd = 0x41
    binsonBeginArray = 0x42
    binsonEndArray = 0x43 
    binsonTrue = 0x44
    binsonFalse = 0x45
    binsonInteger1 = 0x10 
    binsonInteger2 = 0x11 
    binsonInteger4 = 0x12
    binsonInteger8 = 0x13
    binsonDouble = 0x46
    binsonString1 = 0x14
    binsonString2 = 0x15
    binsonString4 = 0x16
    binsonBytes1 = 0x18 
    binsonBytes2 = 0x19 
    binsonBytes4 = 0x1a
)

type field interface {
    toBytes() []byte
}

type Binson map[binsonString]field
type BinsonArray []field
type binsonInt int64
type binsonString string
type binsonBytes []byte
type binsonBool bool
type binsonFloat float64

func packInteger(value int64) []byte {
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
    buf.WriteByte(binsonBegin)
    for _, key := range b.FieldNames(){
        buf.Write(binsonString(key).toBytes())
        buf.Write(b[binsonString(key)].toBytes())
    }
    buf.WriteByte(binsonEnd)
    return buf.Bytes()
}

func (b *BinsonArray) ToBytes() []byte {
    return b.toBytes()
}

func (b *BinsonArray) toBytes() []byte {
    var buf bytes.Buffer
    buf.WriteByte(binsonBeginArray)
    for _, field := range *b {
        buf.Write(field.toBytes())
    }
    buf.WriteByte(binsonEndArray)
    return buf.Bytes()
}

func (a binsonInt) toBytes() []byte {
    buf := new(bytes.Buffer)
    packedInt := packInteger(int64(a))
    switch len(packedInt) {
        case 1:
            buf.WriteByte(binsonInteger1)
        case 2:
            buf.WriteByte(binsonInteger2)
        case 4:
            buf.WriteByte(binsonInteger4)
        case 8:
            buf.WriteByte(binsonInteger8)
        default: 
            panic(fmt.Sprintf("Can not handle byte array of size %d", len(packedInt)))
    }
    buf.Write(packedInt)
    return buf.Bytes()
}

func (a binsonString) toBytes() []byte {
    buf := new(bytes.Buffer)
    lengtBytes := packInteger(int64(len(a)))
    switch len(lengtBytes) {
        case 1:
            buf.WriteByte(binsonString1)
        case 2:
            buf.WriteByte(binsonString2)
        case 4:
            buf.WriteByte(binsonString4)
        default: 
            panic(fmt.Sprintf("Can not handle byte array of size %d", len(lengtBytes)))
    }
    buf.Write(lengtBytes)
    buf.WriteString(string(a))
    return buf.Bytes()
}

func (a binsonBytes) toBytes() []byte {
    buf := new(bytes.Buffer)
    lengtBytes := packInteger(int64(len(a)))
    switch len(lengtBytes) {
        case 1:
            buf.WriteByte(binsonBytes1)
        case 2:
            buf.WriteByte(binsonBytes2)
        case 4:
            buf.WriteByte(binsonBytes4)
        default: 
            panic(fmt.Sprintf("Can not handle byte array of size %d", len(lengtBytes)))
    }
    buf.Write(lengtBytes)
    buf.Write(a)
    return buf.Bytes()
}

func (a binsonBool) toBytes() []byte {
    if a {
        return []byte{binsonTrue}
    } else {
        return []byte{binsonFalse}
    }
}

func (a binsonFloat) toBytes() []byte {
    buf := new(bytes.Buffer)
    buf.WriteByte(binsonDouble)
    binary.Write(buf, binary.LittleEndian, float64(a))
    return buf.Bytes()
}

func readInteger(prefix byte, buf *bytes.Buffer) (int64, error) {
    switch prefix {
        case binsonString1, binsonBytes1, binsonInteger1:
            var value int8 
            err := binary.Read(buf, binary.LittleEndian, &value)
            return int64(value), err

        case binsonString2, binsonBytes2, binsonInteger2:
            var value int16
            err := binary.Read(buf, binary.LittleEndian, &value)
            return int64(value), err

        case binsonString4, binsonBytes4, binsonInteger4:
            var value int32 
            err := binary.Read(buf, binary.LittleEndian, &value)
            return int64(value), err

        case binsonInteger8:
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
        } else if next == binsonEnd {
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
        } else if next == binsonEndArray {
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
        case binsonBegin:
            return parseBinson(buf)
        case binsonBeginArray:
            return parseArray(buf)
        case binsonString1, binsonString2, binsonString4:
            return parseString(start, buf)
        case binsonBytes1, binsonBytes2, binsonBytes4:
            return parseBytes(start, buf)
        case binsonInteger1, binsonInteger2, binsonInteger4, binsonInteger8:
            return parseInteger(start, buf)
        case binsonTrue:
            return true, nil
        case binsonFalse:
            return false, nil
        case binsonDouble:
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