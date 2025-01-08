package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input_small string = `..X...
.SAMX.
.A..A.
XMAS.S
.X....
`

var test_input_large string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestNewBoardSmall(t *testing.T) {
	got := NewBoard(test_input_small)
	want := Board{
		text:   strings.ReplaceAll(test_input_small, "\n", ""),
		width:  6,
		height: 5,
	}
	assert.Equal(t, want, got)
}

func TestNewBoardLarge(t *testing.T) {
	got := NewBoard(test_input_large)
	want := Board{
		text:   strings.ReplaceAll(test_input_large, "\n", ""),
		width:  10,
		height: 10,
	}
	assert.Equal(t, want, got)
}

func TestNewBoardReal(t *testing.T) {
	got := NewBoard(raw_text)
	want := Board{
		text:   strings.ReplaceAll(raw_text, "\n", ""),
		width:  140,
		height: 140,
	}
	assert.Equal(t, want, got)
}

func TestCheckAllDirections(t *testing.T) {
	board := NewBoard(test_input_small)
	tests := []struct {
		name string
		idx  int
		want int
	}{
		{"0", 0, 0}, {"1", 1, 0}, {"2", 2, 1}, {"3", 3, 0}, {"4", 4, 0}, {"5", 5, 0},
		{"6", 6, 0}, {"7", 7, 0}, {"8", 8, 0}, {"9", 9, 0}, {"10", 10, 1}, {"11", 11, 0},
		{"12", 12, 0}, {"13", 13, 0}, {"14", 14, 0}, {"15", 15, 0}, {"16", 16, 0}, {"17", 17, 0},
		{"18", 18, 1}, {"19", 19, 0}, {"20", 20, 0}, {"21", 21, 0}, {"22", 22, 0}, {"23", 23, 0},
		{"24", 24, 0}, {"25", 25, 1}, {"26", 26, 0}, {"27", 27, 0}, {"28", 28, 0}, {"29", 29, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := board.CheckAllDirections(tc.idx, XMAS[:])
			assert.Equal(t, tc.want, got)
		})
	}
}
func TestPart1Small(t *testing.T) {
	got := part1(NewBoard(test_input_small))
	want := 4
	assert.Equal(t, want, got)
}

func TestPart1Large(t *testing.T) {
	got := part1(NewBoard(test_input_large))
	want := 18
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	got := part1(NewBoard(raw_text))
	want := 2613
	assert.Equal(t, want, got)
}

var test2_input string = `M.S
.A.
M.S`

// func TestPart2Small(t *testing.T) {
// 	got := part2(NewBoard(test2_input))
// 	want := 1
// 	assert.Equal(t, want, got)
// }

var test2_large_input string = `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

// func TestPart2Large(t *testing.T) {
// 	got := part2(NewBoard(test2_large_input))
// 	want := 9
// 	assert.Equal(t, want, got)
// }
