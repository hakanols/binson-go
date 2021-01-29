# binson-go
[![Go Reference](https://pkg.go.dev/badge/github.com/hakanols/binson-go.svg)](https://pkg.go.dev/github.com/hakanols/binson-go)

A Golang implementation of Binson. Binson is like JSON, but faster, binary and even simpler. See [binson.org](http://binson.org/).

```
import (
    "fmt"
    "github.com/hakanols/binson-go"
)

func ExampleBinson() {
    b := binson.NewBinson().
        Put("a", 1).
        Put("b", -1).
        Put("c", 250)
    bytes := b.ToBytes()
    fmt.Printf("%X\n", bytes)

    o, _ := binson.Parse(bytes)
    fmt.Printf("HasInt('x'): %t\n", o.HasInt("x"))
    v, _ := o.GetInt("c")
    fmt.Printf("GetInt('c'): %d\n", v)
    // Output: 
    // 40140161100114016210FF14016311FA0041
    // HasInt('x'): false
    // GetInt('c'): 250
}
```
