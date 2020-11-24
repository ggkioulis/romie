package utils

import (
	"errors"
	"os"

	"github.com/romie-gr/romie/internal/exceptions"
)

// EnVar checks if an environment variable exists and returns it.
func EnVar(key string) (string, error) {
	if key == "" {
		return "", &exceptions.ArgError{
			Err: errors.New("received empty key as argument"),
		}
	}

	val, ok := os.LookupEnv(key)

	if val == "" {
		var msg string
		if !ok {
			msg = "environment variable does not exist"
		} else {
			msg = "environment variable has empty value"
		}

		return "", &exceptions.EnvVarError{
			Variable: val,
			Err:      errors.New(msg),
		}
	}

	return val, nil
}
