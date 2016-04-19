package main

import (
        "fmt"
)

type Base struct {
        ID string
}

type Child struct {
        Base
}

func main() {
        child := Child{}
        child.ID = "i am the child"
        fmt.Println(child)         
}

