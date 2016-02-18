package main

import (
	"fmt"

	"github.com/juju/errors"
)

func main() {
	err := fmt.Errorf("%s", "error msg here")
        err = errors.Trace(err)
        err = errors.Annotatef(err, "error has been traces and annotated")

        fmt.Println("details:", errors.Details(err))
	fmt.Println("statck trace:", errors.ErrorStack(err))
}

