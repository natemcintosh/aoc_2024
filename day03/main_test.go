package main

import (
	"testing"

	"github.com/natemcintosh/aoc_2024/utils"
	"github.com/stretchr/testify/assert"
)

var test_input string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestGetMatches(t *testing.T) {
	// Do we get exactly 4 matches from the test input?
	matches := utils.GetGroups(mul_pattern, test_input)
	assert.Equal(t, 4, len(matches))
}

func TestPart1(t *testing.T) {
	got := part1(test_input)
	want := 161
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	got := part1(raw_text)
	want := 162813399
	assert.Equal(t, want, got)
}
