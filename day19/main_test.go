package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestFindMatches(t *testing.T) {
	available_patterns := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	tests := []struct {
		pattern_to_build string
		want_pattern     []string
		want_err         error
	}{
		{"brwrr", []string{"br", "wr", "r"}, nil},
		{"bggr", []string{"b", "g", "g", "r"}, nil},
		{"gbbr", []string{"gb", "br"}, nil},
		{"rrbgbr", []string{"r", "rb", "gb", "r"}, nil},
		{"ubwu", nil, fmt.Errorf("no match found")},
		{"bwurrg", []string{"bwu", "r", "r", "g"}, nil},
		{"brgr", []string{"br", "g", "r"}, nil},
		{"bbrgwb", nil, fmt.Errorf("no match found")},
	}
	for _, tc := range tests {
		got, err := FindMatches(tc.pattern_to_build, available_patterns)
		assert.Equal(t, tc.want_pattern, got)
		if tc.want_err != nil {
			assert.EqualError(t, err, tc.want_err.Error())

		}
	}
}

func TestParseTowels(t *testing.T) {
	building_blocks, desired_patterns := parse_towels(test_input)
	want_building_blocks := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	want_desired_patterns := []string{"brwrr", "bggr", "gbbr", "rrbgbr", "ubwu", "bwurrg", "brgr", "bbrgwb"}

	assert.Equal(t, want_building_blocks, building_blocks)
	assert.Equal(t, want_desired_patterns, desired_patterns)
}

func TestPart1(t *testing.T) {
	building_blocks, desired_patterns := parse_towels(test_input)
	want := 6
	got := part1(desired_patterns, building_blocks)
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	building_blocks, desired_patterns := parse_towels(raw_text)
	want := 276
	got := part1(desired_patterns, building_blocks)
	assert.Equal(t, want, got)
}
