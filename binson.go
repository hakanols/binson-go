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

func (b Binson) HasBinson(name string) bool {
    _, ok := b[binsonString(name)].(Binson)
    return ok
}

func (b Binson) GetBinson(name string) (Binson, bool) {
    obj, ok := b[binsonString(name)].(Binson)
    return obj, ok
}

func (b Binson) HasArray(name string) bool {
    _, ok := b[binsonString(name)].(*BinsonArray)
    return ok
}

func (b Binson) GetArray(name string) (*BinsonArray, bool) {
    obj, ok := b[binsonString(name)].(*BinsonArray)
    return obj, ok
}

func (b Binson) HasInt(name string) bool {
    _, ok := b[binsonString(name)].(binsonInt)
    return ok
}

func (b Binson) GetInt(name string) (int64, bool) {
    obj, ok := b[binsonString(name)].(binsonInt)
    return int64(obj), ok
}

func (b Binson) HasString(name string) bool {
    _, ok := b[binsonString(name)].(binsonString)
    return ok
}

func (b Binson) GetString(name string) (string, bool) {
    obj, ok := b[binsonString(name)].(binsonString)
    return string(obj), ok
}

func (b Binson) HasBytes(name string) bool {
    _, ok := b[binsonString(name)].(binsonBytes)
    return ok
}

func (b Binson) GetBytes(name string) ([]byte, bool) {
    obj, ok := b[binsonString(name)].(binsonBytes)
    return []byte(obj), ok
}

func (b Binson) HasBool(name string) bool {
    _, ok := b[binsonString(name)].(binsonBool)
    return ok
}

func (b Binson) GetBool(name string) (bool, bool) {
    obj, ok := b[binsonString(name)].(binsonBool)
    return bool(obj), ok
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
    key := binsonString(name)
    switch o := value.(type) {
        case Binson:
            b[key] = o
        case *BinsonArray:
            b[key] = o
        case int:
            b[key] = binsonInt(int64(o))
        case int64:
            b[key] = binsonInt(o)
        case string:
            b[key] = binsonString(o)
        case []byte:
            b[key] = binsonBytes(o)
        case bool:
            b[key] = binsonBool(o)
        case float64:
            b[key] = binsonFloat(o)
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
    return b
}