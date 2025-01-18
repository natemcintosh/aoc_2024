package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

// parse_towels takes in the raw input string and returns the list of building blocks
// and the list of desired patterns
func parse_towels(raw string) ([]string, []string) {
	// Split around a blank line
	raw_splits := strings.SplitN(raw, "\n\n", 2)
	raw_bb := raw_splits[0]
	raw_desired := raw_splits[1]

	// Split the building blocks into a slice, splitting on ", "
	// and trimming the trailing newline
	building_blocks := strings.Split(strings.TrimSpace(raw_bb), ", ")

	// Split the desired patterns into a slice, splitting on newlines
	// and trimming the trailing newline
	desired_patterns := strings.Split(strings.TrimSpace(raw_desired), "\n")

	return building_blocks, desired_patterns
}

var NoMatchFound = fmt.Errorf("no match found")

// inner_find_matches takes in a string and a list of patterns and sees which of the available_patterns
// fit at the start of rem. Then it recursively calls itself with the remaining string and the
// remaining patterns.
//
// If it find a complete match, it returns the []string of the patterns that make up the match,
// otherwise it returns an error.
func inner_find_matches(rem string, matched_so_far, available_patterns []string) ([]string, error) {
	// For each match, run find_matches on the remaining string all the
	// available_patterns, passing in the ones matched so far
	for _, m := range available_patterns {
		// If this is not the start of rem, then continue
		if !strings.HasPrefix(rem, m) {
			continue
		}

		// add this m to matched_so_far
		matched_so_far := append(matched_so_far, m)

		// If we've matched all the way to the end, return matched_so_far
		if len(rem) == len(m) {
			return matched_so_far, nil
		}

		// Call the recursive function
		// If it returns an error, continue
		final_match, err := inner_find_matches(rem[len(m):], matched_so_far, available_patterns)

		if err != nil {
			continue
		} else {
			return final_match, nil
		}

	}

	// If we get here, it means we didn't find a match
	return nil, NoMatchFound
}

// FindMatches takes in a string and a list of patterns and sees which of the available_patterns
// fit at the start of rem. Then it recursively calls itself with the remaining string and the
// remaining patterns.
// If it find a complete match, it returns the []string of the patterns that make up the match,
// otherwise it returns an error.
// If it returns an error, it means no match was found.
func FindMatches(to_create string, available_patterns []string) ([]string, error) {
	// Sort the available_patterns by length, longest first
	// This is because we want to try to match the longest patterns first
	slices.SortStableFunc(available_patterns, func(a, b string) int {
		l_diff := len(b) - len(a)
		if l_diff != 0 {
			return l_diff
		}
		return strings.Compare(a, b)
	})

	// Filter available_patterns down to only those that are contained in to_create
	ap := make([]string, 0, len(available_patterns))
	for _, p := range available_patterns {
		if strings.Contains(to_create, p) {
			ap = append(ap, p)
		}
	}

	// If ap is empty, return nil, fmt.Errorf("no match found")
	if len(ap) == 0 {
		return nil, NoMatchFound
	}

	return inner_find_matches(to_create, make([]string, 0), ap)
}

func part1(desired_patterns []string, building_blocks []string) int {
	sum := 0
	for _, dp := range desired_patterns {
		_, err := FindMatches(dp, building_blocks)
		if err == nil {
			sum += 1
		}
	}
	return sum
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Parse Input ===============================================
	parse_start := time.Now()
	building_blocks, desired_patterns := parse_towels(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(desired_patterns, building_blocks)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	// p2 := part2(machines, 101, 103, 100000)
	p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
