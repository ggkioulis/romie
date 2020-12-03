package exceptions

import "testing"

func Test_isWrapped(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWrapped(tt.args.err); got != tt.want {
				t.Errorf("isWrapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareWrappedErrorMessages(t *testing.T) {
	type args struct {
		err               error
		wantSpecificError string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := compareWrappedErrorMessages(tt.args.err, tt.args.wantSpecificError)
			if got != tt.want {
				t.Errorf("compareWrappedErrorMessages() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("compareWrappedErrorMessages() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_compareErrorMessages(t *testing.T) {
	type args struct {
		err               error
		wantSpecificError string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := compareErrorMessages(tt.args.err, tt.args.wantSpecificError)
			if got != tt.want {
				t.Errorf("compareErrorMessages() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("compareErrorMessages() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		err               error
		wantSpecificError string
	}
	tests := []struct {
		name        string
		args        args
		wantMatch   bool
		wantMessage string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatch, gotMessage := Compare(tt.args.err, tt.args.wantSpecificError)
			if gotMatch != tt.wantMatch {
				t.Errorf("Compare() gotMatch = %v, want %v", gotMatch, tt.wantMatch)
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("Compare() gotMessage = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
