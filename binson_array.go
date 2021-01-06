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

func (a *BinsonArray) PutBinson(value Binson) *BinsonArray {
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutInt(value BinsonInt) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutString(value BinsonString) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutBytes(value BinsonBytes) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutBool(value BinsonBool) *BinsonArray{  
    *a = append(*a, value)
    return a
}

func (a *BinsonArray) PutFloat(value BinsonFloat) *BinsonArray {
    *a = append(*a, value)
    return a
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