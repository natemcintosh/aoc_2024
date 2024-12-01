package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

// Node holds the left and right directions from a parent node
type Node struct {
	Left, Right string
}

// Network holds the set of directions (an infinite loop over a set of L/R instructions)
// and a map of nodes. The dir_idx is an index into the directions array; use modulus to
// lookup the next direction.
type Network struct {
	Directions []rune
	dir_idx    uint
	Nodes      map[string]Node
}

// nextDirection returns the next direction in the network's directions array. Will
// loop forever.
func (nw *Network) nextDirection() rune {
	dir := nw.Directions[nw.dir_idx]
	nw.dir_idx = (nw.dir_idx + 1) % uint(len(nw.Directions))
	return dir
}

var nodePattern *regexp.Regexp = regexp.MustCompile(`(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)

// parseInput takes a raw_input string and returns a Network struct
// The input looks like
// ```LLR
//
// AAA = (BBB, BBB)
// BBB = (AAA, ZZZ)
// ZZZ = (ZZZ, ZZZ)```
func parseInput(raw_input string) Network {
	// First, split the raw input by \n\n
	// The first element is the directions
	// The rest are the nodes
	// The nodes are split by \n
	parts := strings.SplitN(raw_input, "\n\n", 2)
	dirs := []rune(parts[0])
	// Check that all the runes are either 'L' or 'R'
	for _, d := range dirs {
		if (d != 'L') && (d != 'R') {
			panic("direction wasn't L or R")
		}
	}

	matches := nodePattern.FindAllStringSubmatch(parts[1], -1)

	// Create a slice of nodes, and fill it
	nodes := make(map[string]Node)
	for _, line := range matches {
		nodes[line[1]] = Node{Left: line[2], Right: line[3]}
	}

	// Create the final Network
	return Network{Directions: dirs, dir_idx: 0, Nodes: nodes}
}

// part1 takes a network and returns the number of nodes visited when starting from AAA
// and ending at ZZZ
func part1(network Network) int {
	count := 0
	curr_key := "AAA"

	// Iterate over the network
	for {
		// Get the next node based on the direction
		switch network.nextDirection() {
		case 'L':
			curr_key = network.Nodes[curr_key].Left
		case 'R':
			curr_key = network.Nodes[curr_key].Right
		default:
			panic("Should have errored when parsing")
		}
		count += 1

		// Check for exit condition
		if curr_key == "ZZZ" {
			break
		}
	}

	return count
}

func main() {
	// Time how long it takes to read the file
	// and parse the games
	parse_start := time.Now()

	// === Parse ====================================================
	raw_text := utils.ReadFile("day08/input.txt")
	nw := parseInput(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(nw)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	// p2_start := time.Now()
	// p2 := part2(games)
	// p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	// fmt.Printf("Part 2 took %v\n", p2_time)

}
