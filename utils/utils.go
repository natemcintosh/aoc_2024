package utils

import (
	"log"
	"os"
)

// ReadFile reads a file and returns its content as a string
func ReadFile(filename string) string {
	byte_contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(byte_contents)
}
