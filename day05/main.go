package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

type Rules struct {
	// Each key is required before its values
	// req_before map[int][]int

	// Each key is required after its values
	req_after map[int][]int

	// Updates are what must be applied to the Rule book
	updates [][]int
}

// NewRules reads text that looks like this
// `47|53
// 97|13
//
// 97,13,61,47,75
// 97,61,53,29,13`
// Where the top half is the rules and the bottom half is the updates
func NewRules(input string) Rules {
	r := Rules{
		// req_before: make(map[int][]int),
		req_after: make(map[int][]int),
		updates:   make([][]int, 0),
	}

	// Split the input into rules and updates
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(parts) != 2 {
		panic("Found more than two parts when splitting by '\\n\\n'")
	}

	// Parse the rules
	rules := strings.Split(parts[0], "\n")
	for _, rule := range rules {
		// Split the rule into the before and after
		parts := strings.Split(rule, "|")
		if len(parts) != 2 {
			panic("More than two parts when splitting by '|'")
		}

		// Parse the before and after
		before := utils.ParseInt(parts[0])
		after := utils.ParseInt(parts[1])

		// Each key is required before its values
		// r.req_before[before] = append(r.req_before[before], after)

		// Each key is required after its values
		r.req_after[after] = append(r.req_after[after], before)
	}

	// Parse the updates
	updates := strings.Split(parts[1], "\n")
	for _, update := range updates {
		// Split the update into the values
		update_line := strings.Split(update, ",")
		if len(update_line) == 0 {
			panic("Could not find any update_line when splitting by ','")
		}

		// Parse the values
		vals := make([]int, 0, len(update_line))
		for _, part := range update_line {
			val := utils.ParseInt(part)
			vals = append(vals, val)
		}

		// Add the update to the list
		r.updates = append(r.updates, vals)
	}

	return r
}

// UpdateIsValid is true if the update is valid according to the rules
// where the r.req_before means that the key must be before the values,
// and the r.req_after means that the key must be after the values
func (r Rules) UpdateIsValid(update []int) bool {
	for idx, val := range update {
		// What must have come before val?
		must_come_before := r.req_after[val]

		// For the items after val
		for _, after := range update[idx+1:] {
			// If the item `before` shows up in `must_come_before`, then it is invalid
			if slices.Contains(must_come_before, after) {
				return false
			}
		}
	}

	return true
}

func part1(rules Rules) int {
	// For each update, check if it is valid according to the rules
	valid_updates := make([][]int, 0)

	for _, update := range rules.updates {
		if rules.UpdateIsValid(update) {
			valid_updates = append(valid_updates, update)
		}
	}

	// Find the middle value of each update that is valid, and sum them
	sum := 0
	for _, update := range valid_updates {
		middle := update[(len(update)-1)/2]
		sum += middle
	}
	return sum
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Parse ====================================================
	// Time how long it takes to read the file and parse the games
	parse_start := time.Now()
	rules := NewRules(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(rules)
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
