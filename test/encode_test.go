package test

import (
	"gotools"
	"testing"
)

func TestBase64EncodeToString(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base64 encode",
			args: args{src: []byte("hello world")},
			want: "aGVsbG8gd29ybGQ=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotools.Base64EncodeToString(tt.args.src); got != tt.want {
				t.Errorf("Base64EncodeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
