package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestUpdateIsValid(t *testing.T) {
	rules := NewRules(test_input)
	tests := []struct {
		name   string
		update []int
		want   bool
	}{
		{"1", rules.updates[0], true},
		{"2", rules.updates[1], true},
		{"3", rules.updates[2], true},
		{"4", rules.updates[3], false},
		{"5", rules.updates[4], false},
		{"6", rules.updates[5], false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := rules.UpdateIsValid(tc.update)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart1(t *testing.T) {
	rules := NewRules(test_input)

	got := part1(rules)
	want := 143
	assert.Equal(t, got, want)
}
