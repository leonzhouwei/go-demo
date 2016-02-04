package main

import (
	_ "github.com/leonzhouwei/go-demo/seelog/log"

	log "github.com/cihub/seelog"
)

func main() {
	log.Tracef("Hello from Seelog!")
}
