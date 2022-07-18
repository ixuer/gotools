package test

import (
	"gotools"
	"testing"
)

func TestFileExists(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "file exists",
			args: args{filepath: "files_test.go"},
			want: true,
		},

		{
			name: "file exists",
			args: args{filepath: "../files.go"},
			want: true,
		},

		{
			name: "file does not exists",
			args: args{filepath: "files_new.go"},
			want: false,
		},

		{
			name: "file does not exists",
			args: args{filepath: "../files_test.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotools.FileExists(tt.args.filepath); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
