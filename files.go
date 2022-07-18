package gotools

import (
	"os"
)

// FileExists returns true if the specified file exists.
func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
