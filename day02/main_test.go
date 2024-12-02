package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const test_input string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

var test_funcs = []tester{is_inc, is_dec, ge1, le3}

func TestPart1(t *testing.T) {
	input := parse(test_input)
	got := part1(input)
	expected := 2
	assert.Equal(t, expected, got)
}

func TestPart1Real(t *testing.T) {
	input := parse(raw_text)
	got := part1(input)
	expected := 663
	assert.Equal(t, expected, got)
}

func TestReport_is_good(t *testing.T) {
	reports := parse(test_input)
	want_vals := []bool{true, false, false, false, false, true}
	tests := []struct {
		name   string
		report []int
		want   bool
	}{
		{"good", reports[0], want_vals[0]},
		{"bad", reports[1], want_vals[1]},
		{"bad", reports[2], want_vals[2]},
		{"bad", reports[3], want_vals[3]},
		{"bad", reports[4], want_vals[4]},
		{"good", reports[5], want_vals[5]},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := report_is_good(tc.report)

			assert.Equal(t, tc.want, got)
		})
	}
}
