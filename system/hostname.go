package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	hostname, error := os.Hostname()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(hostname)
}
