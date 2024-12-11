package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

func create_disk(input string) []int {
	input = strings.TrimSpace(input)
	// Parse each character of the string into an integer
	nums := make([]int, len(input))
	// Convert each rune to an integer by subtracting the rune value of '0'
	// Assumes that all numbers are ascii characters, and I have verified that they are
	for i, r := range input {
		nums[i] = int(r - '0')
	}

	// Each number is the number of spaces taken up by a file or empty space
	total_len := 0
	for _, num := range nums {
		total_len += num
	}

	// Make the slice that will describe the disk space. -1 for empty, and # for file
	// where # is an incrementing integer
	disk := make([]int, total_len)

	// Fill the disk with the file IDs
	file_id := 0
	disk_offset := 0
	for idx, num := range nums {
		// For even idx, this item is a file, and we need to fill the disk with the file_id of length num
		if idx%2 == 0 {
			for i := 0; i < num; i++ {
				disk[disk_offset] = file_id
				disk_offset += 1
			}
			file_id += 1
		} else {
			// For odd idx, this item is a file, and we need to fill the disk with -1 of length num
			for i := 0; i < num; i++ {
				disk[disk_offset] = -1
				disk_offset += 1
			}
		}
	}

	return disk
}

func part1(disk []int) int {
	// The scanner starts at the beginning of the disk
	// The scanner moves one space to the right each time it reads a -1 (empty space)
	// at which point it takes the back most file number and swaps it with the -1
	back_idx := len(disk) - 1

	// Create a copy of the disk
	dc := slices.Clone(disk)
	for idx, block := range dc {
		// If we've reached the index coming backwards, we're done
		if idx >= back_idx {
			break
		}

		if block == -1 {
			// Swap the item at back_idx with the -1
			dc[back_idx], dc[idx] = dc[idx], dc[back_idx]
			// move the back_idx to the next non-empty space
			for dc[back_idx] == -1 {
				back_idx -= 1
			}
		}
	}

	sum := 0
	for idx, block := range dc {
		if block != -1 {
			sum += idx * block
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
	disk := create_disk(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(disk)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	// p2_start := time.Now()
	// p2 := part2(raw_text)
	// p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	// fmt.Printf("Part 2 took %v\n", p2_time)
}
