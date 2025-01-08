package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

type Board struct {
	// The raw text of the board
	text string

	// The width of the board
	width int

	// The height of the board
	height int
}

func NewBoard(raw_input string) Board {
	width := strings.Index(raw_input, "\n")
	if width == -1 {
		panic("Did not find any new line breaks")
	}

	height := strings.Count(raw_input, "\n")

	// We don't want any new line breaks in the text, because they'll mess up the indexing
	ri := strings.ReplaceAll(strings.TrimSpace(raw_input), "\n", "")
	return Board{
		text:   ri,
		width:  width,
		height: height,
	}
}

// The letters that spell "XMAS", but in integer form. Would prefer that this is a const,
// but Go arrays aren't constant.
var XMAS [4]byte = [4]byte{'X', 'M', 'A', 'S'}

// CheckAllDirections checks if the letters in any of the cardinal or diagonal directions.
// Return the number of times `to_find` is spelled in any of the directions.
func (b Board) CheckAllDirections(idx int, to_find []byte) int {
	// Get the row and column indices from the linear index
	row_idx := idx / b.width
	col_idx := idx % b.width
	total_count := 0

	// First, check Up
	if row_idx < 3 {
		// Don't look up
	} else {
		all_good := true
		for i, letter := range to_find {
			li := idx - (b.width * i)
			if b.text[li] != letter {
				all_good = false
				break
			}
		}
		if all_good {
			total_count += 1
		}
	}

	// Check Down
	if row_idx > b.height-4 {
		// Don't look down
	} else {
		all_good := true
		for i, letter := range to_find {
			li := idx + (b.width * i)
			if b.text[li] != letter {
				all_good = false
				break
			}
		}
		if all_good {
			total_count += 1
		}
	}

	// Check Left
	if col_idx < 3 {
		// Don't look left
	} else {
		all_good := true
		for i, letter := range to_find {
			li := idx - i
			if b.text[li] != letter {
				all_good = false
				break
			}
		}
		if all_good {
			total_count += 1
		}

	}

	// Check Right
	if col_idx > b.width-4 {
		// Don't look right
	} else {
		all_good := true
		for i, letter := range to_find {
			li := idx + i
			if b.text[li] != letter {
				all_good = false
				break
			}
		}
		if all_good {
			total_count += 1
		}
	}

	// Check Up-Left
	if row_idx < 3 || col_idx < 3 {
		// Don't look up-left
	} else {
		all_good := true
		for i, letter := range to_find {
			li := idx - (b.width * i) - i
			if b.text[li] != letter {
				all_good = false
				break
			}
		}
		if all_good {
			total_count += 1
		}
	}

	// Check Up-Right
	if row_idx < 3 || col_idx > b.width-4 {
		// Don't look up-right
	} else {
		all_good := true
		for i, letter := range to_find {
			li := idx - (b.width * i) + i
			if b.text[li] != letter {
				all_good = false
				break
			}
		}
		if all_good {
			total_count += 1
		}
	}

	// Check Down-Left
	if row_idx > b.height-4 || col_idx < 3 {
		// Don't look down-left
	} else {
		all_good := true
		for i, letter := range to_find {
			li := idx + (b.width * i) - i
			if b.text[li] != letter {
				all_good = false
				break
			}
		}
		if all_good {
			total_count += 1
		}
	}

	// Check Down-Right
	if row_idx > b.height-4 || col_idx > b.width-4 {
		// Don't look down-right
	} else {
		all_good := true
		for i, letter := range to_find {
			li := idx + (b.width * i) + i
			if b.text[li] != letter {
				all_good = false
				break
			}
		}
		if all_good {
			total_count += 1
		}
	}

	return total_count
}

func part1(board Board) int {
	// Iterate over the board. If it is a "X", check all directions for "XMAS"
	total_count := 0
	for i, letter := range board.text {
		if letter == 'X' {
			total_count += board.CheckAllDirections(i, XMAS[:])
		}
	}
	return total_count
}

// part2 needs to find the pattern
// M.S
// .A.
// M.S
// Or any of its 90 degree rotations
func part2(board Board) int {
	// start_row_idx := 1
	// end_row_idx := board.height - 2
	// start_col_idx := 1
	// end_col_idx := board.width - 2
	return 0
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Parse ====================================================
	// Time how long it takes to read the file and parse the games
	parse_start := time.Now()
	board := NewBoard(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(board)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	// p2_start := time.Now()
	// p2 := part2(board)
	// p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	// fmt.Printf("Part 2 took %v\n", p2_time)
}
