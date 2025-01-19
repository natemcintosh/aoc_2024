package main

import (
	_ "embed"
	"fmt"
	"maps"
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

// inner_find_matches takes in a string and a list of patterns and sees which of the building_blocks
// fit at the start of rem. Then it recursively calls itself with the remaining string and the
// remaining patterns.
//
// If it find a complete match, it adds that to the count of ways the string can be made
func inner_find_matches(rem string, building_blocks []string, n_matches int, n_ways map[string]int) int {
	// If we've already calculated the number of ways to build rem, return it
	if n, ok := n_ways[rem]; ok {
		return n
	}

	// Keep track of the number of ways we can build rem from smaller strings
	n_ways_rem := 0

	// For each match, run find_matches on the remaining string all the
	// available_patterns, passing in the ones matched so far
	for _, m := range building_blocks {
		// If this is not the start of rem, then continue
		if !strings.HasPrefix(rem, m) {
			continue
		}

		// If we've matched all the way to the end, return n_matches + 1
		if len(rem) == len(m) {
			n_matches += 1

			// Add the number of ways we can build rem to the map
			n_ways[rem] += 1
			continue
		}

		// Call the recursive function
		new_ways := inner_find_matches(rem[len(m):], building_blocks, n_matches, n_ways)
		n_matches += new_ways
		n_ways_rem += new_ways
	}

	// Add the number of ways we can build rem to the map
	if n_ways_rem > 0 {
		n_ways[rem] += n_ways_rem
	}

	return n_matches
}

// FindMatches takes in a string and a list of patterns and sees which of the building_blocks
// can be used to build the string. It does this by calling inner_find_matches recursively.
func FindMatches(to_create string, building_blocks []string, n_ways map[string]int) int {
	// Filter available_patterns down to only those that are contained in to_create
	bb := make([]string, 0, len(building_blocks))
	for _, p := range building_blocks {
		if strings.Contains(to_create, p) {
			bb = append(bb, p)
		}
	}

	// If bb is empty, return 0
	if len(bb) == 0 {
		return 0
	}
	inner_find_matches(to_create, bb, 0, n_ways)
	return n_ways[to_create]
}

// prep_map prepares the map of how many ways we can build each pattern. Critically, it
// also checks for any cases where a building block can be built in multiple ways. E.g.
// if we have building blocks ["wr", "w", "r"], then we can build "wr" in two ways: "w" + "r" or "wr".
// Likewise, if we have ["grb", "rb", "gr", "b","r","g"], then we can build "grb" as
// "g" + "rb", "gr" + "b", "grb", or "g" + "r" + "b". The way to deal with this is to
// start with the shortest building blocks and work our way up. This way, we can be sure that
// we include any building blocks that can be built in multiple ways
func prep_map(building_blocks []string) map[string]int {
	// Sort the available_patterns by length, shortest first
	slices.SortStableFunc(building_blocks, func(a, b string) int {
		l_diff := len(a) - len(b)
		if l_diff != 0 {
			return l_diff
		}
		return strings.Compare(a, b)
	})

	n_ways := make(map[string]int, len(building_blocks))

	// Fill with the initial values
	for _, bb := range building_blocks {
		// If the bb is longer than 1 letter, then check how many ways we can build it
		if len(bb) > 1 {
			n_ways[bb] = FindMatches(bb, building_blocks, n_ways)
		} else {
			n_ways[bb] = 1
		}
	}

	return n_ways
}

func solve(desired_patterns []string, building_blocks []string) (int, int) {

	// Create a map for keeping track of the number of ways we can build each pattern
	// This is to avoid recalculating the number of ways to build a pattern
	n_ways := prep_map(building_blocks)

	p1_sum := 0
	p2_sum := 0
	for _, dp := range desired_patterns {
		n := FindMatches(dp, building_blocks, n_ways)
		if n > 0 {
			p1_sum += 1
			p2_sum += n
		}
	}
	// Try summing all the values in n_ways
	p2_v2 := 0
	for v := range maps.Values(n_ways) {
		p2_v2 += v
	}
	return p1_sum, p2_sum
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

	// === Parts 1 and 2 ==============================================
	p1_start := time.Now()
	p1, p2 := solve(desired_patterns, building_blocks)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Parts 1 and 2 took %v\n", p1_time)
}
