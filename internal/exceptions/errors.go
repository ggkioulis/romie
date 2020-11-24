package exceptions

import (
	"fmt"
)

// ArgError denotes that a function's argument was passed incorrectly.
type ArgError struct {
	Err error
}

func (r *ArgError) Error() string {
	return fmt.Sprintf("%T : %v", r, r.Err)
}

// EnvVarError is an error type to track environment variables exceptions.
type EnvVarError struct {
	Variable string
	Err      error
}

func (r *EnvVarError) Error() string {
	return fmt.Sprintf("%T using envar %s : %v", r, r.Variable, r.Err)
}
