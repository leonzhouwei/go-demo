package main

import (
	"fmt"

	"github.com/leonzhouwei/go-demo/go/foo"
)

func main() {
	fmt.Println(foo.Bar)
	foo.Bar = 4242
	fmt.Println(foo.Bar)
}
