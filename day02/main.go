package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

func parse(raw_text string) [][]int {
	lines := strings.Split(strings.TrimSpace(raw_text), "\n")

	res := make([][]int, len(lines))
	for idx, line := range lines {
		nums := strings.Fields(line)
		res[idx] = make([]int, len(nums))
		for i, num := range nums {
			res[idx][i] = utils.ParseInt(num)
		}
	}
	return res
}

// tester is a function that takes an int and checks some condition
type tester func(int) bool

// part1 has the following rules for each report
//
// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func part1(reports [][]int) int {

	count := 0
	// For each report
	for _, r := range reports {
		// Create the diff
		diffs := diff(r)
		if report_is_good(diffs) {
			count += 1
		}
	}
	return count
}

func report_is_good(diffs []int) bool {

	// If the diffs aren't entirely positive or enitrely negative, fail
	if !(all_true(diffs, is_inc) || all_true(diffs, is_dec)) {
		return false
	}

	// If any of the diffs are greater than 3, fail
	if !all_true(diffs, le3) {
		return false
	}

	// If any diffs are less than 1, fail
	if !all_true(diffs, ge1) {
		return false
	}

	return true
}

func diff(x []int) []int {
	res := make([]int, len(x)-1)
	for i := 1; i < len(x); i++ {
		res[i-1] = x[i] - x[i-1]
	}
	return res
}

func all_true(x []int, f tester) bool {
	for _, val := range x {
		if !f(val) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// is_inc is true if the difference is strictly positive
func is_inc(diff int) bool { return diff > 0 }

// is_dec is true if the difference is strictly negative
func is_dec(diff int) bool { return diff < 0 }

// ge1 is true if the difference is greater than or equal to 1
func ge1(diff int) bool { return abs(diff) >= 1 }

// le3 is true if the difference is less than or equal to 3
func le3(diff int) bool { return abs(diff) <= 3 }

func report_is_good_p2(report []int, diffs []int) bool {
	// First check if it is valid via part 1
	if report_is_good(diffs) {
		return true
	}

	// Try removing each element and checking if it is valid
	for idx := range report {
		// Create a new report by removing the element at index i
		new_report := make([]int, len(report)-1)
		copy(new_report, report[:idx])
		copy(new_report[idx:], report[idx+1:])
		new_diffs := diff(new_report)
		if report_is_good(new_diffs) {
			return true
		}
	}
	return false
}

func part2(reports [][]int) int {
	count := 0
	for _, r := range reports {
		diffs := diff(r)
		if report_is_good_p2(r, diffs) {
			count += 1
		}
	}
	return count
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {

	// Time how long it takes to read the file
	// and parse the games
	parse_start := time.Now()

	// === Parse ====================================================
	reports := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(reports)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(reports)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
