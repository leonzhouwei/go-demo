package main

import (
	"fmt"
)

type bird struct {
	name string
}

func (e *bird) setName(name string) {
	e.name = name
}

func (e *bird) hi() {
	fmt.Println(e.name)
}

func main() {
	bird := bird{
		name: "sparrow",
	}
	bird.hi()
	bird.setName("dove")
	bird.hi()
}