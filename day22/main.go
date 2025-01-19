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
	// p2 := part2(secret_numbers)
	p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
