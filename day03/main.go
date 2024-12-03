package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
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

func part2(input string) int {
	// Split the input by `do()`. Assumes input doesn't start with `don't()`. It does
	// not in my input.
	lines := strings.Split(input, "do()")

	sum := 0
	// Run p2_dont_line on each line
	for _, line := range lines {
		sum += part1(strings.Split(line, "don't()")[0])
	}
	return sum
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Part 1 ====================================================
	p1_start := time.Now()
	raw_text = strings.ReplaceAll(strings.TrimSpace(raw_text), "\n", "")
	p1 := part1(raw_text)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(raw_text)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nPart 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
