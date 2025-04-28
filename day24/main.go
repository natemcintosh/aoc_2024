package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/natemcintosh/aoc_2024/circuits"
)

func part1() int {
	// Get the input values
	x, y := circuits.InputValues()

	// Run the circut, and gather the z outputs
	z := circuits.Circuit(x, y)

	// Create a string of ones and zeros, where the earlier index is farther to the
	// right
	slices.Reverse(z)
	var str_z strings.Builder
	for _, bit := range z {
		var ibit int
		if bit {
			ibit = 1
		} else {
			ibit = 0
		}
		str_z.WriteString(strconv.Itoa(ibit))
	}
	z_bits := str_z.String()

	// Intepret each bool as a bit in a binary number. Earliest values
	// are at the lowest bit.
	parsed_val, err := strconv.ParseInt(z_bits, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(parsed_val)
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1()
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	// p2 := part2(graph)
	p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\nPart 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
