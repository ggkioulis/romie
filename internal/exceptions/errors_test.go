package exceptions

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		err     string
		wantErr bool
	}{
		{
			"Convert string to error",
			"This will be converted to an error",
			true,
		},
		{
			"Convert empty string to empty error",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.err); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			} else if err.Error() != tt.err {
				t.Errorf("New() error = %v, wanted %v", err, tt.err)
			}
		})
	}
}
