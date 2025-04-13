package main

import (
	_ "embed"
	"fmt"
	"slices"
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

// DiGraph is a directed graph that stores the edges in both directions. This is
// necessary to be able to traverse the graph in both directions. The graph is
// represented as a map of edges.
type DiGraph struct {
	from_to map[Wire][]Wire
	to_from map[Wire][]Wire
}

func NewDiGraph() DiGraph {
	return DiGraph{
		from_to: make(map[Wire][]Wire),
		to_from: make(map[Wire][]Wire),
	}
}

// AddEdge adds an edge to the graph. It adds the edge in both directions.
func (g *DiGraph) AddEdge(from, to Wire) {
	g.from_to[from] = append(g.from_to[from], to)
	g.to_from[to] = append(g.to_from[to], from)

	// Make sure that there is an entry (even if it is empty) for `to` in the from_to
	// map. This is necessary to be able to traverse the graph in both directions.
	if _, ok := g.from_to[to]; !ok {
		g.from_to[to] = []Wire{}
	}

	// Make sure that there is an entry (even if it is empty) for `from` in the to_from
	// map. This is necessary to be able to traverse the graph in both directions.
	if _, ok := g.to_from[from]; !ok {
		g.to_from[from] = []Wire{}
	}
}

// RemoveEdge removes an edge from the graph. It removes the edge in both directions.
func (g *DiGraph) RemoveEdge(from, to Wire) {
	g.from_to[from] = slices.DeleteFunc(g.from_to[from], func(w Wire) bool { return w == to })
	g.to_from[to] = slices.DeleteFunc(g.to_from[to], func(w Wire) bool { return w == from })
}

// HasEdges checks if there are any edges in the DiGraph at all.
func (g *DiGraph) HasEdges() bool {
	// Check that all from_to slices are not empty
	for _, v := range g.from_to {
		if len(v) > 0 {
			return true
		}
	}

	// Check that all to_from slices are not empty
	for _, v := range g.to_from {
		if len(v) > 0 {
			return true
		}
	}

	return false
}

// TopoSort performs a topological sort on the graph. It returns a slice of wires in
// topological order. It returns an error if the graph is not a DAG (Directed Acyclic
// Graph). This is performed using [Kahn's Algorithm](wikipedia.org/Topological_sorting#Kahn's_algorithm)
func (g *DiGraph) TopoSort() ([]Wire, error) {
	// Create an empty slices that wil container all the sorted elements
	L := make([]Wire, 0, len(g.from_to))

	// Create a slices that we will iterate over. Start with all nodes that have no
	// incoming edges
	S := make([]Wire, 0)

	// Populate S with all nodes that have no incoming edges
	// This is done by checking the to_from map. If a node has no incoming edges, it will
	// have an empty slice in the to_from map.
	for k, v := range g.to_from {
		if len(v) == 0 {
			S = append(S, k)
		}
	}

	// If nothing in S, then fail
	if len(S) == 0 {
		return []Wire{}, fmt.Errorf("no nodes with no incoming edges")
	}

	// While S is not empty, remove a node from S and add it to L
	for len(S) > 0 {
		// Remove a node n from S, and add it to L
		n := S[len(S)-1]
		S = S[:len(S)-1]

		// Add node n to L
		L = append(L, n)

		// For each node m with an edge e from n to m, remove the edge e from the graph
		for idx := 0; idx < len(g.from_to[n]); idx++ {
			m := g.from_to[n][idx]
			g.RemoveEdge(n, m)
			// Decrement the index to account for the removed edge
			idx--

			// If m has not other incoming edges, add it to S
			if len(g.to_from[m]) == 0 {
				S = append(S, m)
			}
		}
	}

	// If the graph still has edges, there there is at least one cycle in the graph
	if g.HasEdges() {
		return []Wire{}, fmt.Errorf("There is a cycle in the graph.")
	}

	return L, nil
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
