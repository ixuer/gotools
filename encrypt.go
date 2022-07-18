package gotools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

// Md5EncryptToString Encrypts the source string to the target string with md5 algorithm.
func Md5EncryptToString(src string) string {
	h := md5.Sum([]byte(src))
	return hex.EncodeToString(h[:])
}

// Md5EncryptToString2 Encrypts the source string to the target string with md5 algorithm.
// It looks much slower than above method.
func Md5EncryptToString2(src string) string {
	w := md5.New()
	_, err := io.WriteString(w, src)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", w.Sum(nil))
}
