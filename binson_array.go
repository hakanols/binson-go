package binson

import (
    "fmt"
)

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

func (a *BinsonArray) HasArray(index int) bool {
	_, ok := (*a)[index].(*BinsonArray)
	return ok
}

func (a *BinsonArray) GetArray(index int) (*BinsonArray, bool) {
    obj, ok := (*a)[index].(*BinsonArray)
	return obj, ok
}

func (a *BinsonArray) PutBinson(value Binson) *BinsonArray {
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) HasBinson(index int) bool {
	_, ok := (*a)[index].(Binson)
	return ok
}

func (a *BinsonArray) GetBinson(index int) (Binson, bool) {
    obj, ok := (*a)[index].(Binson)
	return obj, ok
}

func (a *BinsonArray) PutInt(value BinsonInt) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) HasInt(index int) bool {
	_, ok := (*a)[index].(BinsonInt)
	return ok
}

func (a *BinsonArray) GetInt(index int) (int64, bool) {
    obj, ok := (*a)[index].(BinsonInt)
	return int64(obj), ok
}

func (a *BinsonArray) PutString(value BinsonString) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) HasString(index int) bool {
	_, ok := (*a)[index].(BinsonString)
	return ok
}

func (a *BinsonArray) GetString(index int) (string, bool) {
    obj, ok := (*a)[index].(BinsonString)
	return string(obj), ok
}

func (a *BinsonArray) PutBytes(value BinsonBytes) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) HasBytes(index int) bool {
	_, ok := (*a)[index].(BinsonBytes)
	return ok
}

func (a *BinsonArray) GetBytes(index int) ([]byte, bool) {
    obj, ok := (*a)[index].(BinsonBytes)
	return []byte(obj), ok
}

func (a *BinsonArray) PutBool(value BinsonBool) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) HasBool(index int) bool {
	_, ok := (*a)[index].(BinsonBool)
	return ok
}

func (a *BinsonArray) GetBool(index int) (bool, bool) {
    obj, ok := (*a)[index].(BinsonBool)
	return bool(obj), ok
}

func (a *BinsonArray) PutFloat(value BinsonFloat) *BinsonArray {
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) HasFloat(index int) bool {
	_, ok := (*a)[index].(BinsonFloat)
	return ok
}

func (a *BinsonArray) GetFloat(index int) (float64, bool) {
    obj, ok := (*a)[index].(BinsonFloat)
	return float64(obj), ok
}

func (a *BinsonArray) Put(value interface{}) (*BinsonArray){
    switch o := value.(type) {
        case Binson:
            a.PutBinson(o)
        case *BinsonArray:
            a.PutArray(o)
        case int:
            a.PutInt(BinsonInt(o))
		case BinsonInt:
            a.PutInt(o)	
        case string:
            a.PutString(BinsonString(o))
		case BinsonString:
            a.PutString(o)
        case []byte:
            a.PutBytes(BinsonBytes(o))
		case BinsonBytes:
            a.PutBytes(o)
        case bool:
            a.PutBool(BinsonBool(o))
		case BinsonBool:
            a.PutBool(o)
        case float64:
            a.PutFloat(BinsonFloat(o))
		case BinsonFloat:
            a.PutFloat(o)
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
    return a
}