package utils

import (
	"os"
	"strings"
)

// ReadFile reads a file and returns its content as a string
func ReadFile(filename string) string {
	byte_contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(byte_contents))
}
