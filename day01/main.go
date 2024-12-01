package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

// parse the raw text into two slices of ints, each sorted
func parse(raw_text string) ([]int, []int) {
	lines := strings.Split(raw_text, "\n")

	// Create the output types
	l := make([]int, len(lines))
	r := make([]int, len(lines))

	for idx, line := range lines {
		parts := strings.Fields(line)
		l[idx] = utils.ParseInt(parts[0])
		r[idx] = utils.ParseInt(parts[1])
	}

	// Sort the slices
	slices.Sort(l)
	slices.Sort(r)

	return l, r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(l, r []int) int {
	result := 0
	for idx := range l {
		result += abs(l[idx] - r[idx])
	}
	return result
}

// This time, you'll need to figure out exactly how often each number from the left list
// appears in the right list. Calculate a total similarity score by adding up each
// number in the left list after multiplying it by the number of times that number
// appears in the right list.
func part2(l, r []int) int {
	// Count the number of times each item appears in the right list
	count := make(map[int]int, len(l))

	// Fill the map
	for _, ri := range r {
		// The default value is the zero value, so we can just add to it
		count[ri] += 1
	}

	// For each item in the left, find its frequency, and multiply by it
	res := 0
	for _, li := range l {
		// If it isn't in the map, then it returns the default value (the zero value)
		res += li * count[li]
	}
	return res
}

// Same as above, but use slices.BinarySearch to find the extent of a given number
func part2_v2(l, r []int) int {
	// For each item in the left, find its frequency, and multiply by it
	res := 0
	for _, li := range l {
		// Look up the extent of this item in `r`
		start, _ := slices.BinarySearch(r, li)
		end, _ := slices.BinarySearch(r, li+1)
		n := end - start
		res += li * n
	}

	return res
}

func main() {
	// Time how long it takes to read the file
	// and parse the games
	parse_start := time.Now()

	// === Parse ====================================================
	raw_text := utils.ReadFile("day01/input.txt")
	l, r := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(l, r)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(l, r)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
