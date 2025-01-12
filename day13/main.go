package main

import (
	_ "embed"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

type Button struct {
	Forward_x, Forward_y int
}

type Loc struct {
	X, Y int
}

type ClawMachine struct {
	ButtonA, ButtonB Button
	Prize            Loc
}

type MoveSolution struct {
	NumPushesA, NumPushesB int
}

var NotSolvableError = errors.New("machine is not solvable")

// Cost returns the cost of winning at that machine. If no solution is found, it returns 0
func (c ClawMachine) Cost(skip_over_100 bool) int {
	sol, err := c.FindNumPushes()
	if err != nil {
		return 0
	}
	if (sol.NumPushesA >= 100 || sol.NumPushesB >= 100) && skip_over_100 {
		return 0
	}
	return (sol.NumPushesA * 3) + sol.NumPushesB
}

// FindNumPushes returns the number of times each button was pushed.
// If the machine is not solvable, it returns NotSolvableError.
func (c ClawMachine) FindNumPushes() (MoveSolution, error) {
	na_num := float64(-(c.ButtonB.Forward_x * c.Prize.Y) + c.Prize.X*c.ButtonB.Forward_y)
	na_denom := float64(c.ButtonB.Forward_x*c.ButtonA.Forward_y - c.ButtonA.Forward_x*c.ButtonB.Forward_y)

	if na_denom == 0 {
		return MoveSolution{}, NotSolvableError
	}

	na_res := -na_num / na_denom
	if na_res != float64(int(na_res)) {
		return MoveSolution{}, NotSolvableError
	}

	nb_num := float64(c.ButtonA.Forward_x*c.Prize.Y - c.Prize.X*c.ButtonA.Forward_y)
	nb_denom := float64(c.ButtonB.Forward_x*c.ButtonA.Forward_y - c.ButtonA.Forward_x*c.ButtonB.Forward_y)

	if nb_denom == 0 {
		return MoveSolution{}, NotSolvableError
	}

	nb_res := -nb_num / nb_denom
	if nb_res != float64(int(nb_res)) {
		return MoveSolution{}, NotSolvableError
	}

	return MoveSolution{int(na_res), int(nb_res)}, nil
}

// parse_button_line takes a line like "Button A: X+94, Y+34" and returns a Button.
// It uses regex to parse the line.
func parse_button_line(line string) Button {
	// The regex pattern
	pattern := `Button [A-Z]: X\+(\d+), Y\+(\d+)$`
	re := regexp.MustCompile(pattern)

	// Find the groups
	groups := utils.GetGroups(re, line)

	// Parse the values
	x := utils.ParseInt(groups[0][0])
	y := utils.ParseInt(groups[0][1])

	return Button{x, y}
}

// parse_prize_line takes a line like "Prize: X=8400, Y=5400" and returns a Loc.
func parse_prize_line(line string) Loc {
	// The regex pattern
	pattern := `Prize: X=(\d+), Y=(\d+)$`
	re := regexp.MustCompile(pattern)

	// Find the groups
	groups := utils.GetGroups(re, line)

	// Parse the values
	x := utils.ParseInt(groups[0][0])
	y := utils.ParseInt(groups[0][1])

	return Loc{x, y}
}

// A single machine comes in the form:
// Button A: X+26, Y+66
// Button B: X+67, Y+21
// Prize: X=12748, Y=12176
func parseMachine(raw_machine string) ClawMachine {
	lines := strings.Split(raw_machine, "\n")

	return ClawMachine{
		ButtonA: parse_button_line(lines[0]),
		ButtonB: parse_button_line(lines[1]),
		Prize:   parse_prize_line(lines[2]),
	}
}

// parse takes the raw text and returns a list of ClawMachines. The input look like this:
// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400
//
// Button A: X+26, Y+66
// Button B: X+67, Y+21
// Prize: X=12748, Y=12176
func parse(raw_text string) []ClawMachine {
	raw_machines := strings.Split(strings.TrimSpace(raw_text), "\n\n")

	machines := make([]ClawMachine, len(raw_machines))
	for i, raw_machine := range raw_machines {
		machines[i] = parseMachine(raw_machine)
	}

	return machines
}

func part1(machines []ClawMachine) int {
	sum := 0
	for _, machine := range machines {
		sum += machine.Cost(true)
	}

	return sum
}

func part2(machines []ClawMachine) int {
	sum := 0
	for _, machine := range machines {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000
		sum += machine.Cost(false)
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
	machines := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(machines)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(machines)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
