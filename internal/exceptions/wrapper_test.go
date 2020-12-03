package exceptions

import "testing"

func Test_getCallerInfo(t *testing.T) {
	tests := []struct {
		name         string
		wantFile     string
		wantFunction string
	}{
		{
			"Get file and ",
			"/usr/local/go/src/testing/testing.go",
			"testing.tRunner",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, _, gotFunction := getCallerInfo()
			if gotFile != tt.wantFile {
				t.Errorf("getCallerInfo() got = %v, want %v", gotFile, tt.wantFile)
			}
			if gotFunction != tt.wantFunction {
				t.Errorf("getCallerInfo() got1 = %v, want %v", gotFunction, tt.wantFunction)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err     error
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrap(tt.args.err, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
