package utils

import (
	"os"
	"regexp"
	"slices"
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

// ParseBool converts a string to a bool and panics if it fails
func ParseBool(s string) bool {
	s = strings.TrimSpace(s)
	if s == "1" || strings.ToLower(s) == "true" {
		return true
	}
	if s == "0" || strings.ToLower(s) == "false" {
		return false
	}
	panic("invalid boolean string")
}

// GetGroups is essentially the same as FindAllStringSubmatch, but it returns
// just the groups and not the full match
func GetGroups(re *regexp.Regexp, s string) [][]string {
	matches := re.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		match = slices.Delete(match, 0, 1)
	}
	return matches
}
