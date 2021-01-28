package binson

import (
    "fmt"
    "sort"
)

func NewBinson() Binson {
    b := make(map[binsonString]field)
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
    _, ok := b[binsonString(name)]
    return ok
}

func (b Binson) Remove(name string) {
    delete(b, binsonString(name))
}

func (b Binson) PutBinson(name string, value Binson) Binson {
    b[binsonString(name)] = value
    return b
}

func (b Binson) HasBinson(name string) bool {
    _, ok := b[binsonString(name)].(Binson)
    return ok
}

func (b Binson) GetBinson(name string) (Binson, bool) {
    obj, ok := b[binsonString(name)].(Binson)
    return obj, ok
}

func (b Binson) PutArray(name string, value *BinsonArray) Binson {
    b[binsonString(name)] = value
    return b
}

func (b Binson) HasArray(name string) bool {
    _, ok := b[binsonString(name)].(*BinsonArray)
    return ok
}

func (b Binson) GetArray(name string) (*BinsonArray, bool) {
    obj, ok := b[binsonString(name)].(*BinsonArray)
    return obj, ok
}

func (b Binson) PutInt(name string, value int64) Binson {
    b[binsonString(name)] = binsonInt(value)
    return b
}

func (b Binson) HasInt(name string) bool {
    _, ok := b[binsonString(name)].(binsonInt)
    return ok
}

func (b Binson) GetInt(name string) (int64, bool) {
    obj, ok := b[binsonString(name)].(binsonInt)
    return int64(obj), ok
}

func (b Binson) PutString(name string, value string) Binson {
    b[binsonString(name)] = binsonString(value)
    return b
}

func (b Binson) HasString(name string) bool {
    _, ok := b[binsonString(name)].(binsonString)
    return ok
}

func (b Binson) GetString(name string) (string, bool) {
    obj, ok := b[binsonString(name)].(binsonString)
    return string(obj), ok
}

func (b Binson) PutBytes(name string, value []byte) Binson {
    b[binsonString(name)] = binsonBytes(value)
    return b
}

func (b Binson) HasBytes(name string) bool {
    _, ok := b[binsonString(name)].(binsonBytes)
    return ok
}

func (b Binson) GetBytes(name string) ([]byte, bool) {
    obj, ok := b[binsonString(name)].(binsonBytes)
    return []byte(obj), ok
}

func (b Binson) PutBool(name string, value bool) Binson {
    b[binsonString(name)] = binsonBool(value)
    return b
}

func (b Binson) HasBool(name string) bool {
    _, ok := b[binsonString(name)].(binsonBool)
    return ok
}

func (b Binson) GetBool(name string) (bool, bool) {
    obj, ok := b[binsonString(name)].(binsonBool)
    return bool(obj), ok
}

func (b Binson) PutFloat(name string, value float64) Binson {
    b[binsonString(name)] = binsonFloat(value)
    return b
}

func (b Binson) HasFloat(name string) bool {
    _, ok := b[binsonString(name)].(binsonFloat)
    return ok
}

func (b Binson) GetFloat(name string) (float64, bool) {
    obj, ok := b[binsonString(name)].(binsonFloat)
    return float64(obj), ok
}

func (b Binson) Put(name string, value interface{}) (Binson) {
    switch o := value.(type) {
        case Binson:
            b.PutBinson(name, o)
        case *BinsonArray:
            b.PutArray(name, o)
        case int:
            b.PutInt(name, int64(o))
        case int64:
            b.PutInt(name, o)
        case string:
            b.PutString(name, o)
        case []byte:
            b.PutBytes(name, o)
        case bool:
            b.PutBool(name, o)
        case float64:
            b.PutFloat(name, o)
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
    return b
}