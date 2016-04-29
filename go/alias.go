package main

import (
        "fmt"
)

type AA struct {
        A
}

type A struct {
        Data string
}

func (e *A) Hi() {
        fmt.Println("hi", e.Data)
}

func main() {
        a := &AA{}
        a.Data = "c"
        a.Hi()
}

