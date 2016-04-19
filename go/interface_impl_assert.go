package main

import (
        "fmt"
)

type Hello interface {
        Hi() 
}

type HelloImpl1 struct {
        a string        
}

func (e *HelloImpl1) Hi() {
        fmt.Println("hello from 1")
}

type HelloImpl2 struct {
        a string        
}

func (e *HelloImpl2) Hi() {
        fmt.Println("hello from 2")
}

func main() {
        var hi Hello
        hi = &HelloImpl1{}

        if i, ok := hi.(*HelloImpl1); ok {
                i.Hi()
        } else {
                fmt.Println("not 1")
        }

        if i, ok := hi.(*HelloImpl2); ok {
                i.Hi()
        } else {
                fmt.Println("not 2")
        }
        fmt.Println("oops:", hi.Hi)
}

