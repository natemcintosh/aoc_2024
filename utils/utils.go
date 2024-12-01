package utils

import (
	"os"
	"strconv"
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

// ParseInt converts a string to an int and panics if it fails
func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// ParseFloat converts a string to a float64 and panics if it fails
func ParseFloat(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return n
}
