package utils

import (
	"os"
)

// ReadFile reads a file and returns its content as a string
func ReadFile(filename string) string {
	byte_contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(byte_contents)
}
