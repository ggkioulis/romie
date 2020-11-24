package utils

import (
	"os"
	"reflect"
	"testing"

	"github.com/romie-gr/romie/internal/exceptions"
)

var (
	keyWithValue = "KEY_WITH_VALUE"
	emptyKey     = "EMPTY_KEY"
)

var varValue = map[string]string{
	keyWithValue: "value",
	emptyKey:     "",
}

func setupTest() {
	for key, value := range varValue {
		os.Setenv(key, value)
	}
}

func TestEnVar(t *testing.T) {
	setupTest()

	tests := []struct {
		name    string
		key     string
		want    string
		wantErr error
	}{
		{
			"Get environment variable that exists and has a value",
			keyWithValue,
			varValue[keyWithValue],
			nil,
		},
		{
			"Get environment variable that exists but has empty value",
			emptyKey,
			varValue[emptyKey],
			&exceptions.EnvVarError{},
		},
		{
			"Get environment variable that does not exist",
			"MISSING_KEY",
			"",
			&exceptions.EnvVarError{},
		},
		{
			"Receive empty key as argument",
			"",
			"",
			&exceptions.ArgError{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := EnVar(tt.key)
			if reflect.TypeOf(err) != reflect.TypeOf(tt.wantErr) {
				t.Errorf("EnVar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EnVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
