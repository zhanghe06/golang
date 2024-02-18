package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo() error {
	err := sql.ErrNoRows
	// errors.Wrap with Stack
	return errors.Wrap(err, "foo failed")
}

func bar() error {
	err := foo()
	// errors.WithMessage without Stack
	return errors.WithMessage(err, "bar failed")
}

func main() {
	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err) // format %v print without stack
		fmt.Printf("%+v\n", err) // format %+v print with stack
		return
	}
	if err != nil {
		// unknown error
		fmt.Printf("unknown error")
	}
}
