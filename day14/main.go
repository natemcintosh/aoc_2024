package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

type Robot struct {
	x, y   int
	vx, vy int
}

// parse_robots takes in multiple lines that looks like
// p=0,4 v=3,-3
// p=10,3 v=-1,2
// And returns []Robot.
func parse_robots(raw_text string) []Robot {
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+)\sv=(-?\d+),(-?\d+)`)
	matches := utils.GetGroups(re, raw_text)

	robots := make([]Robot, len(matches))
	for i, m := range matches {
		x, y := utils.ParseInt(m[0]), utils.ParseInt(m[1])
		vx, vy := utils.ParseInt(m[2]), utils.ParseInt(m[3])
		robots[i] = Robot{x, y, vx, vy}
	}

	return robots
}

// PropNSteps will move the robot n steps in the direction of its velocity.
// When it goes over the edge of the map, it will wrap around.
func (r Robot) PropNSteps(n, board_x, board_y int) Robot {
	new_x := (r.x + (n * r.vx)) % board_x
	if new_x < 0 {
		new_x = board_x + new_x
	} else {
		new_x = new_x % board_x
	}

	// In this case, a positive number means the robot is moving down, and vice versa.
	new_y := (r.y + (n * r.vy)) % board_y
	if new_y < 0 {
		new_y = board_y + new_y
	} else {
		new_y = new_y % board_y
	}

	return Robot{
		x:  new_x,
		y:  new_y,
		vx: r.vx,
		vy: r.vy,
	}
}

func CalcSafetyFactor(robots []Robot, n_steps, board_x, board_y int) int {
	rbs := slices.Clone(robots)
	// Propagate each robot n_steps.
	for i, r := range robots {
		rbs[i] = r.PropNSteps(n_steps, board_x, board_y)
	}

	// The four quadrants, start out with 0 in each. The order is NW, NE, SE, SW.
	quadrants := [4]int{0, 0, 0, 0}

	mid_x := board_x / 2
	mid_y := board_y / 2

	for _, r := range rbs {
		// If the robot is on one of the middle lines, ignore it.
		if r.x == mid_x || r.y == mid_y {
			continue
		}

		// Figure out which quadrant it is in.
		if (r.x < mid_x) && (r.y < mid_y) {
			quadrants[0] += 1
		} else if (r.x > mid_x) && (r.y < mid_y) {
			quadrants[1] += 1
		} else if (r.x > mid_x) && (r.y > mid_y) {
			quadrants[2] += 1
		} else {
			quadrants[3] += 1
		}
	}

	// Calculate the product of the quadrants.
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func PrintBoard(robots []Robot, board_x, board_y int) {
	board := make([][]int, board_y)
	for i := range board {
		board[i] = make([]int, board_x)
	}

	for _, r := range robots {
		board[r.y][r.x] += 1
	}

	for _, row := range board {
		for _, cell := range row {
			if cell > 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

// longest_nonzero_consecutive finds the longest consective non-zero sequence of
// identical values in a slice.
func longest_nonzero_consecutive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// Initialize the current value and count
	curr := arr[0]
	curr_count := 1
	max_count := 1

	for i := 1; i < len(arr); i++ {
		// If it's zero, reset the count and skip
		if arr[i] == 0 {
			curr = 0
			curr_count = 0
			continue
		}

		if arr[i] == curr {
			curr_count += 1
		} else {
			curr = arr[i]
			curr_count = 1
		}

		if curr_count > max_count {
			max_count = curr_count
		}
	}

	return max_count
}

func part2(robots []Robot, board_x, board_y, max_iters int) int {
	// Make a 2D slice for the board.
	arr := make([][]int, board_y)
	for i := range arr {
		arr[i] = make([]int, board_x)
	}

	for step_n := 0; step_n < max_iters; step_n++ {
		// Make sure we start with an empty map.
		for _, row := range arr {
			clear(row)
		}

		// Propagate all the robots one step
		for j, r := range robots {
			nr := r.PropNSteps(1, board_x, board_y)
			robots[j] = nr
			arr[nr.y][nr.x] += 1
		}

		// Look for 10+ consecutive columns with at least one robot in each.
		for _, row := range arr {
			if longest_nonzero_consecutive(row) >= 8 {
				PrintBoard(robots, board_x, board_y)
				return step_n
			}
		}
	}
	return -1
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Parse Input ===============================================
	parse_start := time.Now()
	machines := parse_robots(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := CalcSafetyFactor(machines, 100, 101, 103)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(machines, 101, 103, 100000)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
