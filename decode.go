package gotools

import "encoding/base64"

// Base64Decode Decode string to bytes slice.
func Base64Decode(src string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Base64DecodeToString Decode base64 string to string.
func Base64DecodeToString(src string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
