package main

import (
	"strings"
	//	"fmt"

	"github.com/leonzhouwei/go-demo/hdfs/client"

	"github.com/qiniu/log.v1"
)

func init() {
	log.SetOutputLevel(log.Ldebug)
}

func main() {
	config := client.Config{
		Username: "zhouwei",
	}
	c := client.NewHTTPClient(config)

	// write
	writeErr := c.WriteAndClose("/user/zhouwei/output/0606-2.txt", strings.NewReader("oops"))
	if writeErr != nil {
		log.Fatal(writeErr)
	}

	//
	//	homeDir, err := c.GetHomeDir()
	//	if err != nil {
	//		log.Fatal(homeDir)
	//	}
	//	fmt.Println(homeDir)

	//
	//	lsDir, err := c.LsDir("/user/zhouwei")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(lsDir)

}
