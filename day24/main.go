package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

// Wire is the three letters that make up the name of a node.
type Wire [3]rune

// NewWire creates a Wire from a string exactly 3 letters long. All letters must be valid
// ASCII
func NewWire(s string) Wire {
	if len(s) != 3 {
		panic("ID must be exactly 3 characters long")
	}

	var wire Wire
	for idx, r := range s {
		if r > 127 {
			panic("ID must be valid ASCII")
		}
		wire[idx] = r
	}

	return wire
}

// String shows the three letters in an easy to read manner
func (wire Wire) String() string {
	return fmt.Sprintf("%c%c%c", wire[0], wire[1], wire[2])
}

func (wire Wire) Format(f fmt.State, c rune) {
	switch c {
	case 's':
		fmt.Fprintf(f, "%s", wire.String())
	case 'v':
		fmt.Fprintf(f, "%v", wire.String())
	default:
		fmt.Fprintf(f, "%s", wire.String())
	}
}

// NameFromInt fetches the original Wire that created this int64. Assumes that all wires
// are created only from valid ASCII characters.
func NameFromInt(n int64) Wire {
	return Wire{
		rune((n >> 32) & 0xFF),
		rune((n >> 16) & 0xFF),
		rune(n & 0xFF),
	}
}

// ID merges the three runes into a single int64. Because the runes are all valid ASCII,
// and we keep them shifted 16 places apart, they won't overlap and mess eachother up.
// It should in theory not produce hash collisions for this case.
func (wire Wire) ID() int64 {
	return (int64(wire[0]) << 32) | (int64(wire[1]) << 16) | int64(wire[2])
}

type op int

const (
	AND op = iota
	OR
	XOR
)

func NewOp(s string) op {
	switch s {
	case "AND":
		return AND
	case "OR":
		return OR
	case "XOR":
		return XOR
	default:
		panic("Invalid operation")
	}
}

type Gate struct {
	// The input wires
	in1, in2 Wire

	// The output wire
	out Wire

	// The operation to perform
	op op
}

// NewGate creates a new gate from a string. Below are examples of the strings that can
// be passed in:
// ntg XOR fgs -> mjb
// y02 OR x01 -> tnw
// kwq OR kpj -> z05
// x00 OR x03 -> fst
// tgd XOR rvg -> z01
// vdt OR tnw -> bfw
// bfw AND frj -> z10
// ffh OR nrd -> bqk
func NewGate(s string) Gate {
	// Split around spaces
	parts := strings.Split(s, " ")
	return Gate{
		NewWire(parts[0]),
		NewWire(parts[2]),
		NewWire(parts[4]),
		NewOp(parts[1]),
	}
}

// Calc takes two input values, calcultes the output of the gate, and returns the
// output value.
func (g Gate) Calc(in1, in2 bool) bool {
	switch g.op {
	case AND:
		return in1 && in2
	case OR:
		return in1 || in2
	case XOR:
		return in1 != in2
	default:
		panic("Invalid operation")
	}
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Parse Input ===============================================
	parse_start := time.Now()
	// graph := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	// p1 := part1(graph)
	p1_time := time.Since(p1_start)
	// fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	// p2 := part2(graph)
	p2_time := time.Since(p2_start)
	// fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
