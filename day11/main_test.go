package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	tests := []struct {
		name   string
		stone  int
		in_map map[int]int
		want   map[int]int
	}{
		{"0", 0, map[int]int{0: 1}, map[int]int{1: 1}},
		{"0", 0, map[int]int{0: 2}, map[int]int{1: 2}},
		{"0", 0, map[int]int{0: 10}, map[int]int{1: 10}},
		{"1", 1, map[int]int{1: 1}, map[int]int{2024: 1}},
		{"1", 1, map[int]int{1: 2}, map[int]int{2024: 2}},
		{"1", 1, map[int]int{1: 10}, map[int]int{2024: 10}},
		{"10", 10, map[int]int{10: 1}, map[int]int{1: 1, 0: 1}},
		{"10", 10, map[int]int{10: 2}, map[int]int{1: 2, 0: 2}},
		{"10", 10, map[int]int{10: 10}, map[int]int{1: 10, 0: 10}},
		{"99", 99, map[int]int{99: 1}, map[int]int{9: 2}},
		{"99", 99, map[int]int{99: 2}, map[int]int{9: 4}},
		{"99", 99, map[int]int{99: 10}, map[int]int{9: 20}},
		{"999", 999, map[int]int{999: 1}, map[int]int{2021976: 1}},
		{"999", 999, map[int]int{999: 2}, map[int]int{2021976: 2}},
		{"999", 999, map[int]int{999: 10}, map[int]int{2021976: 10}},
		{"72", 72, map[int]int{72: 1}, map[int]int{7: 1, 2: 1}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := update_stone(tc.stone, tc.in_map, make(map[int]int))
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSolve(t *testing.T) {
	stones := parse("125 17")
	tests := []struct {
		name  string
		steps int
		want  int
	}{
		{"1", 1, 3}, {"2", 2, 4}, {"3", 3, 5}, {"4", 4, 9},
		{"5", 5, 13}, {"6", 6, 22}, {"25", 25, 55312},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := solve(stones, tc.steps)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParts1And2(t *testing.T) {
	stones := parse("6563348 67 395 0 6 4425 89567 739318")
	tests := []struct {
		name  string
		steps int
		want  int
	}{
		{"25", 25, 184927}, {"75", 75, 220357186726677},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := solve(stones, tc.steps)
			assert.Equal(t, tc.want, got)
		})
	}
}
