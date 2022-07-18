package test

import (
	"gotools"
	"reflect"
	"testing"
)

func TestBase64Decode(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "base64 encode",
			args:    args{src: "aGVsbG8gd29ybGQ="},
			want:    []byte("hello world"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gotools.Base64Decode(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base64Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64DecodeToString(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "base64 decode to string",
			args:    args{src: "aGVsbG8gd29ybGQ="},
			want:    "hello world",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gotools.Base64DecodeToString(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64DecodeToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Base64DecodeToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
