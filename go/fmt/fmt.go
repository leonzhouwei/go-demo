package main

import (
	"fmt"
)

func main() {
	var f float64 = float64(0.1)
	str := fmt.Sprintln("key", "=", f)
	fmt.Println(str)
}
