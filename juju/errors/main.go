package main

import (
	"fmt"

	"github.com/juju/errors"
)

func main() {
	err := fmt.Errorf("%s", "error msg here")
	err2 := fmt.Errorf("err2: %s", err)
	fmt.Println(err2)
        err = errors.Trace(err)
        err = errors.Annotatef(err, "error has been traces and annotated")
	
	fmt.Println(err)
        fmt.Println("details:", errors.Details(err))
	fmt.Println("stack trace:", errors.ErrorStack(err))
}

