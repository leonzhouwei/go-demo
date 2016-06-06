package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

func main() {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	resp, err := kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		//		fmt.Println("oops:", err)
		msg := fmt.Sprintf("%s", err)
		if strings.HasPrefix(msg, "100:") {
			fmt.Println("oops:", msg)
		}
		_, err2 := kapi.Set(context.Background(), "foo", "bar", nil)
		if err != nil {
			log.Fatal(err2)
		}
	} else {
		fmt.Println(resp.Node.Value)
	}
}
