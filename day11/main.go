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

// Applies each rule in order, quiting if one matches. The `storage` has the stone number
// as the key, and the number of stones with that number as the value.
func update_stone(stone_val int, storage map[int]int) map[int]int {
	// Rule 1: if the value is 0, set it to 1
	if stone_val == 0 {
		// Get the number of stones with this value
		storage[1] += storage[0]
		delete(storage, 0)
		return storage
	}

	// Rule 2: if the value has an even number of digits, it is split in two, where the
	// first half is the first half of the digits and the second half is the second half
	// of the digits
	s := strconv.Itoa(stone_val)
	if len(s)%2 == 0 {
		mid := len(s) / 2
		first_half, _ := strconv.Atoi(s[:mid])
		second_half, _ := strconv.Atoi(s[mid:])
		storage[first_half] += storage[stone_val]
		storage[second_half] += storage[stone_val]
		delete(storage, stone_val)
		return storage
	}

	// Rule 3: if none of the above, multiply the number by 2024
	storage[stone_val*2024] += storage[stone_val]
	delete(storage, stone_val)
	return storage
}

func solve(stones []int, n_steps int) int {
	// Create a map for holding the number of stones at a given number
	sm := make(map[int]int, len(stones))

	// Populate the map with the initial stones
	for _, stone := range stones {
		sm[stone] += 1
	}

	// For n_steps, update all the stones in `sm`
	for i := 0; i < n_steps; i++ {
		// Make a copy of sm to put the new values in
		sm_copy := maps.Clone(sm)
		// For each stone in `sm`
		for stone_val := range sm {
			sm_copy = update_stone(stone_val, sm_copy)
		}

		// Overwrite sm with sm_copy
		sm = maps.Clone(sm_copy)
		fmt.Printf("After step %d: %v\n", i+1, sm)
	}

	// Count how many of each type of stone we have
	sum := 0
	for _, num := range sm {
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
	// p2 := solve(stones, 26)
	p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
