package errors

import (
	"errors"
	"log"
)

func DieIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func New(message string) error {
	return errors.New(message)
}
