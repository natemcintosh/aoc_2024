package main

import (
	"testing"

	"github.com/natemcintosh/aoc_2024/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	raw_input := `3   4
4   3
2   5
1   3
3   9
3   3`
	l, r := parse(raw_input)
	got := part1(l, r)
	want := 11
	assert.Equal(t, want, got)
}

func TestPart1RealInput(t *testing.T) {
	raw_text := utils.ReadFile("input.txt")
	l, r := parse(raw_text)
	got := part1(l, r)
	want := 1646452
	assert.Equal(t, want, got)
}

func TestPart2(t *testing.T) {
	raw_input := `3   4
4   3
2   5
1   3
3   9
3   3`
	l, r := parse(raw_input)
	got := part2(l, r)
	want := 31
	assert.Equal(t, want, got)
}

func TestPart2RealInput(t *testing.T) {
	raw_text := utils.ReadFile("input.txt")
	l, r := parse(raw_text)
	got := part2(l, r)
	want := 23609874
	assert.Equal(t, want, got)
}
