package test

import (
	"gotools"
	"testing"
)

func TestMd5EncryptToString(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "md5 encrypt",
			args: args{src: "hello world"},
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
		{
			name: "md5 encrypt",
			args: args{src: "abc"},
			want: "900150983cd24fb0d6963f7d28e17f72",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotools.Md5EncryptToString(tt.args.src); got != tt.want {
				t.Errorf("Md5Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5Encrypt2(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "md5 encrypt",
			args: args{src: "hello world"},
			want: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},

		{
			name: "md5 encrypt",
			args: args{src: "abc"},
			want: "900150983cd24fb0d6963f7d28e17f72",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gotools.Md5EncryptToString2(tt.args.src); got != tt.want {
				t.Errorf("Md5Encrypt2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMd5Encrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotools.Md5EncryptToString("hello world")
	}
}

/*
goos: darwin
goarch: amd64
pkg: gotools
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMd5Encrypt
BenchmarkMd5Encrypt-12    	 6554540	       169.1 ns/op
PASS
*/

func BenchmarkMd5Encrypt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotools.Md5EncryptToString2("hello world")
	}
}

/*
goos: darwin
goarch: amd64
pkg: gotools
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMd5Encrypt2
BenchmarkMd5Encrypt2-12    	 3349638	       337.2 ns/op
PASS
*/
