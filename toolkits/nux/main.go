package main

import (
	"fmt"
	"log"

	"github.com/toolkits/nux"
)

func main() {
	printDiskUsage1()
	printDiskUsage2()
}

func printDiskUsage1() {
	path := "/tmp"
	diskUsage, err := disk.Usage(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(diskUsage)
}

func printDiskUsage2() {
	path := "/"
	diskUsage, err := disk.Usage(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(diskUsage)
}
