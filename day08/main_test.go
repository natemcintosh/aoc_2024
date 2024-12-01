package main

import (
	"testing"

	"github.com/natemcintosh/aoc_2024/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	raw_input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
	network := parseInput(raw_input)
	got := part1(network)
	want := 6
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	raw_input := utils.ReadFile("input.txt")
	network := parseInput(raw_input)
	got := part1(network)
	want := 12643
	assert.Equal(t, want, got)
}

func TestParseInput(t *testing.T) {
	raw_input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
	got := parseInput(raw_input)
	want := Network{
		Directions: []rune("LLR"),
		dir_idx:    0,
		Nodes: map[string]Node{
			"AAA": {Left: "BBB", Right: "BBB"},
			"BBB": {Left: "AAA", Right: "ZZZ"},
			"ZZZ": {Left: "ZZZ", Right: "ZZZ"},
		}}
	assert.Equal(t, want, got)
}

func TestNextDirection(t *testing.T) {
	network := Network{
		Directions: []rune("LR"),
		dir_idx:    0,
	}
	dirs := []rune("LR")
	for i := 0; i < 100; i++ {
		got := network.nextDirection()
		want := dirs[i%2]
		assert.Equal(t, want, got)
	}
}
