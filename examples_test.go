package binson_test

import (
    "fmt"
    "math"
    "github.com/hakanols/binson-go"
)

func ExampleNewBinson() {
    b := binson.NewBinson()
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 4041
}

func ExampleNewPutInt() {
    b := binson.NewBinson().
        Put("cid", 4)
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 401403636964100441
}

func ExamplePutBinson() {
    b := binson.NewBinson().
        PutBinson("a", binson.NewBinson().
            PutInt("b", 2))
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 401401614014016210024141
}

func ExamplePut_Nested() {
    b := binson.NewBinson().
        Put("a", 1).
        Put("b", binson.NewBinson().
            Put("c", 3)).
        Put("d", 4)
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 40140161100114016240140163100341140164100441
}

func ExampleNewBinsonArray() {
    b := binson.NewBinson().
        PutArray("a", binson.NewBinsonArray().
            PutInt(1).
            PutString("hello"))
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 40140161421001140568656C6C6F4341
}

func ExamplePutArray() {
    b := binson.NewBinson().
        Put("a", 1).
        Put("b", binson.NewBinsonArray().
            Put(10).
            Put(20)).
        Put("c", 3)
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 40140161100114016242100A101443140163100341
}

func ExamplePutArray_Nested() {
    b := binson.NewBinson().
        Put("a", 1).
        Put("b", binson.NewBinsonArray().
            Put(10).
            Put(binson.NewBinsonArray().
                Put(100).
                Put(101)).
            Put(20)).
        Put("c", 3)
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 40140161100114016242100A421064106543101443140163100341
}

func ExamplePut() {
    b := binson.NewBinson().
        Put("a", 1).
        Put("b", -1).
        Put("c", 250).
        Put("d", math.MaxInt32).
        Put("f", math.MaxInt64)
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 40140161100114016210FF14016311FA0014016412FFFFFF7F14016613FFFFFFFFFFFFFF7F41
}

func ExampleToBytes() {
    b := binson.NewBinson().
        Put("aaaa", 250)
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 4014046161616111FA0041
}

func ExamplePutString() {
    b := binson.NewBinson().
        PutString("aaaa", "bbb")
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 40140461616161140362626241
}

func ExamplePutBytes() {
    b := binson.NewBinson().
        PutBytes("aa", []byte{5, 5, 5})
    fmt.Printf("%X\n", b.ToBytes())
    // Output: 4014026161180305050541
}