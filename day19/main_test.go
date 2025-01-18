package main

import (
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

func TestPrepMap(t *testing.T) {
	building_blocks := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	want := map[string]int{
		"b":   1,
		"r":   1,
		"g":   1,
		"wr":  1,
		"rb":  2,
		"gb":  2,
		"br":  2,
		"bwu": 1,
	}
	got := prep_map(building_blocks)
	assert.Equal(t, want, got)
}

func TestFindMatches(t *testing.T) {
	building_blocks := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	tests := []struct {
		pattern_to_build string
		want_count       int
	}{
		{"brwrr", 2},
		{"bggr", 1},
		{"gbbr", 4},
		{"rrbgbr", 6},
		{"ubwu", 0},
		{"bwurrg", 1},
		{"brgr", 2},
		{"bbrgwb", 0},
	}
	for _, tc := range tests {
		got := FindMatches(tc.pattern_to_build, building_blocks, make(map[string]int))
		assert.Equal(t, tc.want_count, got)
	}
}

func TestParseTowels(t *testing.T) {
	building_blocks, desired_patterns := parse_towels(test_input)
	want_building_blocks := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	want_desired_patterns := []string{"brwrr", "bggr", "gbbr", "rrbgbr", "ubwu", "bwurrg", "brgr", "bbrgwb"}

	assert.Equal(t, want_building_blocks, building_blocks)
	assert.Equal(t, want_desired_patterns, desired_patterns)
}

func TestPart1And2(t *testing.T) {
	building_blocks, desired_patterns := parse_towels(test_input)
	p1_want := 6
	p2_want := 16
	p1_got, p2_got := solve(desired_patterns, building_blocks)
	assert.Equal(t, p1_want, p1_got)
	assert.Equal(t, p2_want, p2_got)
}

func TestPart1And2Real(t *testing.T) {
	building_blocks, desired_patterns := parse_towels(raw_text)
	p1_want := 276
	// p2_want
	p1_got, _ := solve(desired_patterns, building_blocks)
	assert.Equal(t, p1_want, p1_got)
}
