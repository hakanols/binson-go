package binson

import (
    "fmt"
)

func NewBinsonArray() *BinsonArray {
    a := BinsonArray([]field{})
    return &a
}

func (a *BinsonArray) Size() int{
    return len(*a);
}

func (a *BinsonArray) Remove(index int){
    *a = append( (*a)[:index], (*a)[index+1:]...)
}

func (a *BinsonArray) inRange(index int) bool{
    return index < 0 || a.Size() <= index
}

func (a *BinsonArray) HasArray(index int) bool {
    if a.inRange(index){
        return false
    }
    _, ok := (*a)[index].(*BinsonArray)
    return ok
}

func (a *BinsonArray) GetArray(index int) (*BinsonArray, bool) {
    obj, ok := (*a)[index].(*BinsonArray)
    return obj, ok
}

func (a *BinsonArray) HasBinson(index int) bool {
    if a.inRange(index){
        return false
    }
    _, ok := (*a)[index].(Binson)
    return ok
}

func (a *BinsonArray) GetBinson(index int) (Binson, bool) {
    obj, ok := (*a)[index].(Binson)
    return obj, ok
}

func (a *BinsonArray) HasInt(index int) bool {
    if a.inRange(index){
        return false
    }
    _, ok := (*a)[index].(binsonInt)
    return ok
}

func (a *BinsonArray) GetInt(index int) (int64, bool) {
    obj, ok := (*a)[index].(binsonInt)
    return int64(obj), ok
}

func (a *BinsonArray) HasString(index int) bool {
    if a.inRange(index){
        return false
    }
    _, ok := (*a)[index].(binsonString)
    return ok
}

func (a *BinsonArray) GetString(index int) (string, bool) {
    obj, ok := (*a)[index].(binsonString)
    return string(obj), ok
}

func (a *BinsonArray) HasBytes(index int) bool {
    if a.inRange(index){
        return false
    }
    _, ok := (*a)[index].(binsonBytes)
    return ok
}

func (a *BinsonArray) GetBytes(index int) ([]byte, bool) {
    obj, ok := (*a)[index].(binsonBytes)
    return []byte(obj), ok
}

func (a *BinsonArray) HasBool(index int) bool {
    if a.inRange(index){
        return false
    }
    _, ok := (*a)[index].(binsonBool)
    return ok
}

func (a *BinsonArray) GetBool(index int) (bool, bool) {
    obj, ok := (*a)[index].(binsonBool)
    return bool(obj), ok
}

func (a *BinsonArray) HasFloat(index int) bool {
    if a.inRange(index){
        return false
    }
    _, ok := (*a)[index].(binsonFloat)
    return ok
}

func (a *BinsonArray) GetFloat(index int) (float64, bool) {
    obj, ok := (*a)[index].(binsonFloat)
    return float64(obj), ok
}

func (a *BinsonArray) addField(value field) *BinsonArray {
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) Put(value interface{}) (*BinsonArray){
    switch o := value.(type) {
        case Binson:
			a.addField(o)
        case *BinsonArray:
			a.addField(o)
        case int:
            a.addField(binsonInt(int64(o)))
        case int64:
            a.addField(binsonInt(o))
        case string:
            a.addField(binsonString(o))
        case []byte:
            a.addField(binsonBytes(o))
        case bool:
            a.addField(binsonBool(o))
        case float64:
            a.addField(binsonFloat(o))
			
        default: 
            panic(fmt.Sprintf("%T is not handeled by Binson", o))
    }
    return a
}