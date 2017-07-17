package main

import (
	"fmt"
	"errors"
)

// newErr combines a custom comment and error into a new formatted error.
func newErr(input string, err error) error {
	if err != nil {
		return errors.New(input + " " + err.Error())
	}
	return errors.New(input)
}

// ComboErr combines a custom comment and two error messages into one
// formatted error.
func ComboErr(input string, first error, second error) error {
	var res string
	switch {
	case first != nil && second != nil:
		res = fmt.Sprintf("%s %s. %s", input, first.Error(),
			second.Error())
		return errors.New(res)
	case first != nil:
		return newErr(input, first)
	case second != nil:
		return newErr(input, second)
	}
	return errors.New(input)
}