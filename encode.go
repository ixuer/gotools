package gotools

import "encoding/base64"

// Base64EncodeToString Encode source bytes slice to target string.
func Base64EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
