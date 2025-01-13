package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input string = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestParseRobots(t *testing.T) {
	want := []Robot{
		{0, 4, 3, -3},
		{6, 3, -1, -3},
		{10, 3, -1, 2},
		{2, 0, 2, -1},
		{0, 0, 1, 3},
		{3, 0, -2, -2},
		{7, 6, -1, -3},
		{3, 0, -1, -2},
		{9, 3, 2, 3},
		{7, 3, -1, 2},
		{2, 4, 2, -3},
		{9, 5, -3, -3},
	}
	got := parse_robots(test_input)
	assert.Equal(t, want, got)
}

func TestPropNSteps(t *testing.T) {
	r := Robot{2, 4, 2, -3}

	tests := []struct {
		name string
		n    int
		want Robot
	}{
		{"1 steps", 1, Robot{4, 1, 2, -3}},
		{"2 steps", 2, Robot{6, 5, 2, -3}},
		{"3 steps", 3, Robot{8, 2, 2, -3}},
		{"4 steps", 4, Robot{10, 6, 2, -3}},
		{"5 steps", 5, Robot{1, 3, 2, -3}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := r.PropNSteps(tc.n, 11, 7)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart1(t *testing.T) {
	robots := parse_robots(test_input)
	got := CalcSafetyFactor(robots, 100, 11, 7)
	want := 12
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	robots := parse_robots(raw_text)
	got := CalcSafetyFactor(robots, 100, 101, 103)
	want := 228410028
	assert.Equal(t, want, got)
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"1", []int{1, 2, 3}, 1},
		{"2", []int{1, 1, 2, 3}, 2}, {"2", []int{1, 1, 2, 2, 3, 3}, 2},
		{"3", []int{1, 1, 1}, 3}, {"0", []int{}, 0},
		{"3-non-zero", []int{0, 0, 0, 0, 1, 1, 1}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := longest_nonzero_consecutive(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart2Real(t *testing.T) {
	robots := parse_robots(raw_text)
	got := part2(robots, 101, 103, 10000)
	want := 8258
	assert.Equal(t, want, got)
}
