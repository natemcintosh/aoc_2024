package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

func parse(raw_text string) []int {
	// Split on newlines, and parse each line as an int
	lines := strings.Split(strings.TrimSpace(raw_text), "\n")

	secret_numbers := make([]int, len(lines))
	for i, line := range lines {
		secret_numbers[i] = utils.ParseInt(line)
	}
	return secret_numbers
}

// mix will calculate the bitwise XOR of the given value and the secret number. Then,
// the secret number becomes the result of that operation. (If the secret number is 42
// and you were to mix 15 into the secret number, the secret number would become 37.)
func mix(secret, given int) int {
	return secret ^ given
}

// prune calculates the value of the secret number modulo 16777216. Then, the secret
// number becomes the result of that operation. (If the secret number is 100000000 and
// you were to prune the secret number, the secret number would become 16113920.)
func prune(secret int) int {
	return secret % 16777216
}

// ones_place returns the ones place of the given number
func ones_place(secret int) int {
	return secret % 10
}

// step follows these three steps to get the next pseudorandom number:
//  1. Calculate the result of multiplying the secret number by 64. Then, mix this result
//     into the secret number. Finally, prune the secret number.
//  2. Calculate the result of dividing the secret number by 32. Round the result down
//     to the nearest integer. Then, mix this result into the secret number. Finally, prune
//     the secret number.
//  3. Calculate the result of multiplying the secret number by 2048. Then, mix this
//     result into the secret number. Finally, prune the secret number.
func step(secret int) int {
	// Step 1
	secret = prune(mix(secret*64, secret))

	// Step 2
	secret = prune(mix(secret/32, secret))

	// Step 3
	secret = prune(mix(secret*2048, secret))

	return secret
}

// ChangeTracker keeps track of the differences of 5 values. However, we only need the
// i and (i - 1) values, and then we can keep track of the 4 sets of differences.
//
// It is not suggested to create this struct directly. Instead, always use NewChangeTracker
type ChangeTracker struct {
	// The values. New values move "in from the right", and old values move "out to the left"
	vals [2]int

	// Shift all the values to the left, and make the new right-most item
	// diffs[3] = vals[1] - vals[0]
	diffs [4]int
}

func NewChangeTracker(vals [2]int) ChangeTracker {
	diffs := [4]int{0, 0, 0, vals[1] - vals[0]}
	return ChangeTracker{vals, diffs}
}

// Push will add a new value to the end of the list, and remove the first element. It
// will also update the differences.
func (ct *ChangeTracker) Push(value int) {
	ct.vals[0] = ct.vals[1]
	ct.vals[1] = value

	ct.diffs[0] = ct.diffs[1]
	ct.diffs[1] = ct.diffs[2]
	ct.diffs[2] = ct.diffs[3]
	ct.diffs[3] = ct.vals[1] - ct.vals[0]
}

// track_all_changes_for_seller will go through 2000 steps for a seller's starting
// number, and keep track of the most recent 4 changes in ones place value of the secret
// number. For set of 4 changes, it will store the value of the ones place. If a set
// of 4 changes has already been seen, it will ignore this value.
// If a set of 4 changes has not been seen, it will store the value of the ones place.
//
// Finally, it will return a map of the values of the ones place for each set of 4 changes
func track_all_changes_for_seller(secret, n_steps int, changes map[[4]int]int) map[[4]int]int {
	// We re-use the allocations in changes to avoid allocations. So clear the values here
	clear(changes)
	ct := NewChangeTracker([2]int{ones_place(secret), ones_place(step(secret))})
	secret = step_n(secret, 2)

	for idx := 2; idx < n_steps; idx++ {
		// Get the ones place of the secret number
		ones := ones_place(secret)

		// Add the ones place to the change tracker
		ct.Push(ones)

		// If the diffs have been filled with at least 4 values, then start checking
		if idx >= 4 {
			if _, ok := changes[ct.diffs]; !ok {
				changes[ct.diffs] = ones
			}
		}

		// Step the secret number
		secret = step(secret)
	}
	return changes
}

// step_n repeats the step function n times
func step_n(secret int, n int) int {
	for i := 0; i < n; i++ {
		secret = step(secret)
	}
	return secret
}

// Sum the values of the secret numbers after 2000 steps
func part1(secrets []int) int {
	sum := 0
	for _, secret := range secrets {
		sum += step_n(secret, 2000)
	}
	return sum
}

// merge_maps will merge two maps together, summing the values of any duplicate keys.
// Note that this will modify the first map in place.
func merge_maps(m1, m2 map[[4]int]int) map[[4]int]int {
	for k, v := range m2 {
		if existingValue, exists := m1[k]; exists {
			m1[k] = existingValue + v
		} else {
			m1[k] = v
		}
	}
	return m1
}

func part2(secrets []int) int {
	// Track all the changes for each secret number
	changes := make(map[[4]int]int, len(secrets))
	changes2 := make(map[[4]int]int, len(secrets))

	for _, secret := range secrets {
		changes = merge_maps(
			changes,
			track_all_changes_for_seller(secret, 2000, changes2),
		)
	}

	// Return the max value found
	max_val := 0
	for _, val := range changes {
		if val > max_val {
			max_val = val
		}
	}

	return max_val
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Parse Input ===============================================
	parse_start := time.Now()
	secret_numbers := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(secret_numbers)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(secret_numbers)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
