package e

import (
	"fmt"
	"log"
)

// Print error msg and error.
func Wrap(msg string, e error) error {
	err := fmt.Errorf("%s: %w", msg, e)
	log.Println(err)
	return err
}

// Check error msg and error.
func WrapIfErr(msg string, e error) error {
	if e == nil {
		return nil
	}

	return Wrap(msg, e)
}
