package main

import (
	_ "embed"
	"fmt"
	"maps"
	"strconv"
	"strings"
	"time"
)

func parse(raw_input string) []int {
	var result []int
	for _, v := range strings.Fields(strings.Trim(raw_input, "\n ")) {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}
	return result
}

// Applies each rule in order, quiting if one matches. The `in_map` has the stone number
// as the key, and the number of stones with that number as the value. The `out_map` is
// the map that will be updated and returned.
func update_stone(stone_val int, in_map, out_map map[int]int) map[int]int {
	// Rule 1: if the value is 0, set it to 1
	if stone_val == 0 {
		// Get the number of stones with this value
		out_map[1] += in_map[0]
		return out_map
	}

	// Rule 2: if the value has an even number of digits, it is split in two, where the
	// first half is the first half of the digits and the second half is the second half
	// of the digits
	s := strconv.Itoa(stone_val)
	if len(s)%2 == 0 {
		mid := len(s) / 2
		first_half, _ := strconv.Atoi(s[:mid])
		second_half, _ := strconv.Atoi(s[mid:])
		out_map[first_half] += in_map[stone_val]
		out_map[second_half] += in_map[stone_val]
		return out_map
	}

	// Rule 3: if none of the above, multiply the number by 2024
	out_map[stone_val*2024] += in_map[stone_val]
	return out_map
}

func solve(stones []int, n_steps int) int {
	// Create a map for holding the number of stones at a given number
	in_map := make(map[int]int, len(stones))

	// Populate the map with the initial stones
	for _, stone := range stones {
		in_map[stone] += 1
	}

	// Create the out_map
	out_map := make(map[int]int)

	// For n_steps, update all the stones in `sm`
	for range n_steps {
		// Make sure we start with a fresh out_map each iteration
		clear(out_map)

		// For each stone in `in_map`
		for stone_val := range in_map {
			out_map = update_stone(stone_val, in_map, out_map)
		}

		// Overwrite in_map with out_map
		in_map = maps.Clone(out_map)
	}

	// Count how many of each type of stone we have
	sum := 0
	for _, num := range in_map {
		sum += num
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
	stones := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := solve(stones, 25)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := solve(stones, 75)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
