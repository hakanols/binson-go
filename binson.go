package binson

import (
    "fmt"
    "sort"
)

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

func (b Binson) HasBinson(name BinsonString) bool {
	_, ok := b[name].(Binson)
	return ok
}

func (b Binson) GetBinson(name BinsonString) (Binson, bool) {
    obj, ok := b[name].(Binson)
	return obj, ok
}

func (b Binson) PutArray(name BinsonString, value *BinsonArray) Binson {
    b[name] = value
    return b
}

func (b Binson) HasArray(name BinsonString) bool {
	_, ok := b[name].(*BinsonArray)
	return ok
}

func (b Binson) GetArray(name BinsonString) (*BinsonArray, bool) {
    obj, ok := b[name].(*BinsonArray)
	return obj, ok
}

func (b Binson) PutInt(name BinsonString, value BinsonInt) Binson {
    b[name] = value
    return b
}

func (b Binson) HasInt(name BinsonString) bool {
	_, ok := b[name].(BinsonInt)
	return ok
}

func (b Binson) GetInt(name BinsonString) (int64, bool) {
    obj, ok := b[name].(BinsonInt)
	return int64(obj), ok
}

func (b Binson) PutString(name BinsonString, value BinsonString) Binson {
    b[name] = value
    return b
}

func (b Binson) HasString(name BinsonString) bool {
	_, ok := b[name].(BinsonString)
	return ok
}

func (b Binson) GetString(name BinsonString) (string, bool) {
    obj, ok := b[name].(BinsonString)
	return string(obj), ok
}

func (b Binson) PutBytes(name BinsonString, value BinsonBytes) Binson {
    b[name] = value
    return b
}

func (b Binson) HasBytes(name BinsonString) bool {
	_, ok := b[name].(BinsonBytes)
	return ok
}

func (b Binson) GetBytes(name BinsonString) ([]byte, bool) {
    obj, ok := b[name].(BinsonBytes)
	return []byte(obj), ok
}

func (b Binson) PutBool(name BinsonString, value BinsonBool) Binson {
    b[name] = value
    return b
}

func (b Binson) HasBool(name BinsonString) bool {
	_, ok := b[name].(BinsonBool)
	return ok
}

func (b Binson) GetBool(name BinsonString) (bool, bool) {
    obj, ok := b[name].(BinsonBool)
	return bool(obj), ok
}

func (b Binson) PutFloat(name BinsonString, value BinsonFloat) Binson {
    b[name] = value
    return b
}

func (b Binson) HasFloat(name BinsonString) bool {
	_, ok := b[name].(BinsonFloat)
	return ok
}

func (b Binson) GetFloat(name BinsonString) (float64, bool) {
    obj, ok := b[name].(BinsonFloat)
	return float64(obj), ok
}

func (b Binson) Put(name BinsonString, value interface{}) (Binson) {
    switch o := value.(type) {
        case Binson:
            b.PutBinson(name, o)
        case *BinsonArray:
            b.PutArray(name, o)
        case int:
            b.PutInt(name, BinsonInt(o))
		case BinsonInt:
            b.PutInt(name, o)
        case string:
            b.PutString(name, BinsonString(o))
		case BinsonString:
            b.PutString(name, o)
        case []byte:
            b.PutBytes(name, BinsonBytes(o))
		case BinsonBytes:
            b.PutBytes(name, o)
        case bool:
            b.PutBool(name, BinsonBool(o))
		case BinsonBool:
            b.PutBool(name, o)
		case float64:
            b.PutFloat(name, BinsonFloat(o))
		case BinsonFloat:
            b.PutFloat(name, o)
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
    return b
}