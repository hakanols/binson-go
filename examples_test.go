package binson_test

import (
    "fmt"
    "github.com/hakanols/binson-go"
)


func ExampleNewBinson() {
    b := binson.NewBinson().
	    Put("cid", 4);
	fmt.Printf("%X\n", b.ToBytes())
    // Output: 401403636964100441
}