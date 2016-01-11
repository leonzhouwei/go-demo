package main

import (
	"fmt"
	"log"
	"syscall"
)

// DiskStatus disk capacity status
type DiskStatus struct {
	All  uint64 // count of all btyes of a disk
	Used uint64 // count of used btyes of a disk
	Free uint64 // count of free btyes of a disk
}

func main() {
	ds, err := diskUsageUnixLike(".")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(ds)
		fmt.Println(" all:", ds.All/1024/1024/1024, "GB")
		usedfloat64 := float64(ds.Used)
		allfloat64 := float64(ds.All)
		usedRatio := usedfloat64 / allfloat64
		fmt.Printf("used: %.2f%%, %v GB\n", usedRatio*100.0, ds.Used/1024/1024/1024)
		fmt.Printf("free: %.2f%%, %v GB\n", (1.0-usedRatio)*100.0, ds.Free/1024/1024/1024)
	}
}

func diskUsageUnixLike(path string) (*DiskStatus, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return nil, err
	}
	ds := new(DiskStatus)
	ds.All = fs.Blocks * uint64(fs.Bsize)
	ds.Free = fs.Bfree * uint64(fs.Bsize)
	ds.Used = ds.All - ds.Free
	return ds, nil
}
