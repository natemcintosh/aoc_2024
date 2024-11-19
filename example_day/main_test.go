package main

import (
	"strings"
	"testing"

	"github.com/natemcintosh/aoc_2024/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseGame(t *testing.T) {
	// Create a table of test cases
	tests := []struct {
		name     string
		input    string
		expected []RGB
	}{
		{
			name:  "Game 1",
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: []RGB{
				{B: 3, R: 4},
				{R: 1, G: 2, B: 6},
				{G: 2},
			},
		},
		{
			name:  "Game 2",
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected: []RGB{
				{B: 1, G: 2},
				{G: 3, B: 4, R: 1},
				{G: 1, B: 1},
			},
		},
		{
			name:  "Game 3",
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: []RGB{
				{G: 8, B: 6, R: 20},
				{B: 5, R: 4, G: 13},
				{G: 5, R: 1},
			},
		},
		{
			name:  "Game 4",
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected: []RGB{
				{G: 1, R: 3, B: 6},
				{G: 3, R: 6},
				{G: 3, B: 15, R: 14},
			},
		},
		{
			name:  "Game 5",
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected: []RGB{
				{R: 6, B: 1, G: 3},
				{B: 2, R: 1, G: 2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := parse_game(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPart1(t *testing.T) {
	raw_input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	lines := strings.Split(raw_input, "\n")
	games := make([]Draws, len(lines))
	for i, line := range lines {
		games[i] = parse_game(line)
	}
	actual := part1(games)
	assert.Equal(t, 8, actual)
}

func TestPart1RealInput(t *testing.T) {
	raw_text := utils.ReadFile("input.txt")
	lines := strings.Split(raw_text, "\n")
	games := make([]Draws, len(lines))
	for i, line := range lines {
		games[i] = parse_game(line)
	}
	got := part1(games)

	assert.Equal(t, 2449, got)
}

func TestPart2(t *testing.T) {
	raw_input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	lines := strings.Split(raw_input, "\n")
	games := make([]Draws, len(lines))
	for i, line := range lines {
		games[i] = parse_game(line)
	}

	got := part2(games)
	assert.Equal(t, 2286, got)
}

func TestPart2Real(t *testing.T) {
	raw_text := utils.ReadFile("input.txt")
	lines := strings.Split(raw_text, "\n")
	games := make([]Draws, len(lines))
	for i, line := range lines {
		games[i] = parse_game(line)
	}
	got := part2(games)

	assert.Equal(t, 63981, got)
}
