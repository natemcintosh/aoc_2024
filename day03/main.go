package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

var mul_pattern *regexp.Regexp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func part1(input string) int {
	matches := utils.GetGroups(mul_pattern, input)
	// Parse each group into digits, multiply them, and sum the results
	sum := 0
	for _, match := range matches {
		a := utils.ParseInt(match[0])
		b := utils.ParseInt(match[1])
		sum += a * b
	}

	return sum
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// // === Parse ====================================================
	// // Time how long it takes to read the file and parse the games
	// parse_start := time.Now()
	// reports := parse(raw_text)
	// parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(raw_text)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	// p2_start := time.Now()
	// p2 := part2(reports)
	// p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("Part 1 took %v\n", p1_time)
	// fmt.Printf("Part 2 took %v\n", p2_time)
}
