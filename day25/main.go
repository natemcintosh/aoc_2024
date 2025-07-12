package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

// LockKey represents the five numbers that make up a lock or a key. They are always
// between 0 and 5 inclusive.
type LockKey struct {
	// vals represents the five numbers that make up a lock or a key. They are always
	// between 0 and 5 inclusive.
	vals [5]uint8
}

// fits checks that there is no overlap between the two LockKeys
func (lk LockKey) fits(other LockKey) bool {
	if (lk.vals[0]+other.vals[0] <= 5) &&
		(lk.vals[1]+other.vals[1] <= 5) &&
		(lk.vals[2]+other.vals[2] <= 5) &&
		(lk.vals[3]+other.vals[3] <= 5) &&
		(lk.vals[4]+other.vals[4] <= 5) {
		return true
	}
	return false
}

// parse_lockey parses a string representing a lock or a key into a LockKey struct. If
// the first line is all dots (`.....`) then it is a key, otherwise it is a lock. The
// first and last lines do not contribute to the count. If it is a key, the returned
// bool is true, otherwise false.
func parse_lockey(raw_text string) (lk LockKey, is_lock bool) {
	// Split on newlines to get each row
	rows := strings.Split(raw_text, "\n")

	// Is this a lock or key?
	if rows[0] == "....." {
		is_lock = false
	} else {
		is_lock = true
	}

	// Parse the values from the rows. For each row, add to the column sums.
	for i, row := range rows {
		if i == 0 || i == len(rows)-1 {
			continue
		}
		for j, char := range row {
			if char == '#' {
				lk.vals[j] += 1
			}
		}
	}

	return lk, is_lock
}

// parse reads the raw text, and converts it into two slices of locks and keys.
func parse(raw_text string) (locks, keys []LockKey) {
	// Split on double new lines
	lock_keys := strings.SplitSeq(raw_text, "\n\n")

	for lock_key := range lock_keys {
		lk, is_lock := parse_lockey(lock_key)
		if is_lock {
			locks = append(locks, lk)
		} else {
			keys = append(keys, lk)
		}
	}

	return
}

func part1(locks, keys []LockKey) int {
	sum := 0
	for _, l := range locks {
		for _, k := range keys {
			if l.fits(k) {
				sum += 1
			}
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
	locks, keys := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(locks, keys)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
}
