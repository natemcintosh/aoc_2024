package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

// A compressed disk entry
type DiskEntry struct {
	// If for empty space, file_id is -1
	file_id int

	// The start index of this entry
	start int

	// How long is this entry
	length int
}

// CheckSum calculates the "dot product" of the indices of the entry, with the file_id
func (de DiskEntry) CheckSum() int {
	// If it's an empty file, it has checksum of 0
	if de.file_id == -1 {
		return 0
	}

	// Sum of numbers between a and b, where a <= b is (b-a+1)(a+b)/2
	b := de.start + de.length - 1
	a := de.start
	return de.file_id * (b - a + 1) * (a + b) / 2
}

func create_disk(input string) ([]int, []DiskEntry) {
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
	compressed_disk := make([]DiskEntry, 0)

	// Fill the disk with the file IDs
	file_id := 0
	disk_offset := 0
	for idx, num := range nums {
		// For even idx, this item is a file, and we need to fill the disk with
		// the file_id of length num
		if idx%2 == 0 {
			// Add the compressed disk entry, if it has length greater than 0
			if num > 0 {
				compressed_disk = append(compressed_disk, DiskEntry{
					file_id: file_id,
					start:   disk_offset,
					length:  num,
				})
			}

			// Add the file_id to the disk
			for i := 0; i < num; i++ {
				disk[disk_offset] = file_id
				disk_offset += 1
			}
			file_id += 1
		} else {
			// Add the compressed disk entry, if it has length greater than 0
			if num > 0 {
				compressed_disk = append(compressed_disk, DiskEntry{
					file_id: -1,
					start:   disk_offset,
					length:  num,
				})
			}

			// For odd idx, this item is a file, and we need to fill the disk with
			// -1 of length num
			for i := 0; i < num; i++ {
				disk[disk_offset] = -1
				disk_offset += 1
			}
		}
	}

	return disk, compressed_disk
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

// This time, attempt to move whole files to the leftmost span of free space blocks that
// could fit the file. Attempt to move each file *exactly once* in order of decreasing
// file ID number starting with the file with the highest file ID number. If there is no
// span of free space to the left of a file that is large enough to fit the file, the
// file does not move.
func part2(compressed_disk []DiskEntry) int {
	// Create a copy of the disk
	disk := slices.Clone(compressed_disk)

	// Sort the disk by start point
	slices.SortFunc(disk, func(a, b DiskEntry) int {
		return a.start - b.start
	})

	// Keep track of which file ids we've seen.
	seen := make(map[int]struct{})

	// Iterate through the disk in reverse order
	for right_idx := len(disk) - 1; right_idx >= 0; right_idx-- {
		file_to_swap := disk[right_idx]
		// fmt.Printf("Considering file %+v\n", file_to_swap)

		// If we've already seen this one, or it is empty, skip
		if file_to_swap.file_id == -1 {
			// fmt.Printf("Skipping empty file\n")
			continue
		} else if _, ok := seen[file_to_swap.file_id]; ok {
			// fmt.Printf("Skipping already seen file\n")
			continue
		} else {
			// Add this file id to the ones we've seen
			// fmt.Printf("Adding %+v to seen: %+v\n", file_to_swap.file_id, seen)
			seen[file_to_swap.file_id] = struct{}{}
		}

		// For each index leading up to disk_entry_idx
		for left_idx, file_entry := range disk[:right_idx] {
			// If it is not an empty space, continue
			if file_entry.file_id != -1 {
				continue
			}

			// If the empty space is not large enough, continue
			if file_entry.length < file_to_swap.length {
				// fmt.Printf("Skipping file %+v because it is too small\n", file_entry)
				continue
			}

			// fmt.Printf("Inserting %+v into %+v\n", file_to_swap, file_entry)
			// Calculate how much empty space will be left of this empty space once
			// we put the file_to_swap in
			remaining_space := file_entry.length - file_to_swap.length
			// fmt.Printf("There will %+v remaining space after insertion\n", remaining_space)

			// If no empty space is left, simply swap the two entries
			if remaining_space == 0 {
				disk[left_idx], disk[right_idx] = disk[right_idx], disk[left_idx]
				// Swap their start points
				disk[left_idx].start, disk[right_idx].start = disk[right_idx].start, disk[left_idx].start
				// fmt.Printf("Swapped %+v and %+v\n", disk[left_idx], disk[right_idx])
			} else {
				// Insert empty space at right_idx
				disk[right_idx].file_id = -1

				// Insert the file_to_swap into the space
				disk = slices.Insert(disk, left_idx, file_to_swap)
				disk[left_idx].start = file_entry.start
				// fmt.Printf("Disk is now %+v\n", disk)

				// We're inserting a new item, so we have to increment the disk_entry_idx
				right_idx += 1

				// Give the empty space its new start and length
				new_empty_start_idx := file_entry.start + file_to_swap.length
				disk[left_idx+1] = DiskEntry{file_id: -1, start: new_empty_start_idx, length: remaining_space}

			}

			// Break out since we have moved this file
			break
		}
	}

	sum := 0
	for _, de := range disk {
		sum += de.CheckSum()
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
	disk, compressed_disk := create_disk(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(disk)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(compressed_disk)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
