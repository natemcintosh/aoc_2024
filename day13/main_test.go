package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseButtonLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected Button
	}{
		{"one", "Button A: X+94, Y+34", Button{94, 34}},
		{"two", "Button B: X+22, Y+67", Button{22, 67}},
		{"three", "Button B: X+26, Y+47", Button{26, 47}},
		{"four", "Button A: X+13, Y+56", Button{13, 56}},
		{"five", "Button B: X+67, Y+31", Button{67, 31}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			button := parse_button_line(tc.line)
			assert.Equal(t, tc.expected, button)
		})
	}
}

func TestParsePrizeLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected Loc
	}{
		{"one", "Prize: X=8400, Y=5400", Loc{8400, 5400}},
		{"two", "Prize: X=12748, Y=12176", Loc{12748, 12176}},
		{"three", "Prize: X=7870, Y=6450", Loc{7870, 6450}},
		{"four", "Prize: X=18641, Y=10279", Loc{18641, 10279}},
		{"five", "Prize: X=10843, Y=10358", Loc{10843, 10358}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prize := parse_prize_line(tc.line)
			assert.Equal(t, tc.expected, prize)
		})
	}
}

func TestParseMachine(t *testing.T) {
	tests := []struct {
		name        string
		raw_machine string
		expected    ClawMachine
	}{
		{"one", "Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400", ClawMachine{
			ButtonA: Button{94, 34},
			ButtonB: Button{22, 67},
			Prize:   Loc{8400, 5400},
		}},
		{"two", "Button A: X+26, Y+66\nButton B: X+67, Y+21\nPrize: X=12748, Y=12176", ClawMachine{
			ButtonA: Button{26, 66},
			ButtonB: Button{67, 21},
			Prize:   Loc{12748, 12176},
		}},
		{"three", "Button A: X+17, Y+86\nButton B: X+84, Y+37\nPrize: X=7870, Y=6450", ClawMachine{
			ButtonA: Button{17, 86},
			ButtonB: Button{84, 37},
			Prize:   Loc{7870, 6450},
		}},
		{"four", "Button A: X+69, Y+23\nButton B: X+27, Y+71\nPrize: X=18641, Y=10279", ClawMachine{
			ButtonA: Button{69, 23},
			ButtonB: Button{27, 71},
			Prize:   Loc{18641, 10279},
		}},
		{"five", "Button A: X+13, Y+56\nButton B: X+67, Y+31\nPrize: X=10843, Y=10358", ClawMachine{
			ButtonA: Button{13, 56},
			ButtonB: Button{67, 31},
			Prize:   Loc{10843, 10358},
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			machine := parseMachine(tc.raw_machine)
			assert.Equal(t, tc.expected, machine)
		})
	}
}

func TestFindNumPushes(t *testing.T) {
	tests := []struct {
		name         string
		claw         ClawMachine
		expected_sol MoveSolution
		expected_err error
	}{
		{"one", ClawMachine{
			ButtonA: Button{94, 34},
			ButtonB: Button{22, 67},
			Prize:   Loc{8400, 5400},
		}, MoveSolution{80, 40}, nil},
		{"two", ClawMachine{
			ButtonA: Button{26, 66},
			ButtonB: Button{67, 21},
			Prize:   Loc{12748, 12176},
		}, MoveSolution{}, NotSolvableError},
		{"three", ClawMachine{
			ButtonA: Button{17, 86},
			ButtonB: Button{84, 37},
			Prize:   Loc{7870, 6450},
		}, MoveSolution{38, 86}, nil},
		{"four", ClawMachine{
			ButtonA: Button{69, 23},
			ButtonB: Button{27, 71},
			Prize:   Loc{18641, 10279},
		}, MoveSolution{}, NotSolvableError},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			solution, err := tc.claw.FindNumPushes()
			assert.Equal(t, tc.expected_sol, solution)
			assert.Equal(t, tc.expected_err, err)
		})
	}
}

func TestPart1Real(t *testing.T) {
	machines := parse(raw_text)
	want := 31552
	got := part1(machines)
	assert.Equal(t, want, got)
}

func TestPart2Real(t *testing.T) {
	machines := parse(raw_text)
	want := 95273925552482
	got := part2(machines)
	assert.Equal(t, want, got)
}
