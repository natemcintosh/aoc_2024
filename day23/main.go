package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

// Node represents a node in a graph. It's made up of two runes
type Node [2]rune

var BadNodeInput = fmt.Errorf("Could not produce a Node\n")

func (n Node) Format(f fmt.State, c rune) {
	switch c {
	case 'v': // default format
		fmt.Fprintf(f, "Node(%c%c)", n[0], n[1])
	case 's': // string format
		fmt.Fprintf(f, "%c%c", n[0], n[1])
	case 'q': // quoted string format
		fmt.Fprintf(f, "\"%c%c\"", n[0], n[1])
	default:
		fmt.Fprintf(f, "Node(%c%c)", n[0], n[1])
	}
}

// NewNode takes a string that must be exactly two letters, and reates a Node
func NewNode(s string) (Node, error) {
	if len(s) != 2 {
		fmt.Printf("s: %v\n", s)
		return Node{}, BadNodeInput
	}
	var n Node
	for idx, char := range s {
		n[idx] = char
	}

	return n, nil
}

// Graph is an undirected graph. It stores both the nodes, and the edges between nodes
type Graph struct {
	Edges map[[2]Node]struct{}
	Nodes map[Node]struct{}
}

func (g Graph) Format(f fmt.State, c rune) {
	switch c {
	case 'v': // default format
		fmt.Fprintf(f, "Graph{\n  Nodes: %v,\n  Edges: %v\n}", g.Nodes, g.Edges)
	case 's': // compact format
		fmt.Fprintf(f, "Graph(Nodes: %d, Edges: %d)", len(g.Nodes), len(g.Edges))
	case 'q': // quoted JSON-like format
		fmt.Fprintf(f, "\"Graph with %d nodes and %d edges\"", len(g.Nodes), len(g.Edges))
	default:
		fmt.Fprintf(f, "Graph(Nodes: %d, Edges: %d)", len(g.Nodes), len(g.Edges))
	}
}

// parse takes the input text and parses it into a graph
func parse(raw_text string) Graph {
	// Create the empty graph
	g := Graph{make(map[[2]Node]struct{}), make(map[Node]struct{})}
	// For each line, for each pair of letters, add an edge to the graph
	lines := strings.Split(raw_text, "\n")
	for _, line := range lines {
		// Split around the "-"
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic("Could not parse line " + line)
		}

		// Create a new node for each part
		n1, err := NewNode(parts[0])
		if err != nil {
			panic(err)
		}
		n2, err := NewNode(parts[1])
		if err != nil {
			panic(err)
		}
		// Add the edge to the graph
		edge := CreateEdge(n1, n2)

		// Add the edge to the graph
		g.Edges[edge] = struct{}{}

		// Add the nodes to the graph
		g.Nodes[n1] = struct{}{}
		g.Nodes[n2] = struct{}{}

	}
	return g
}

// CreateEdge creates an edge between two nodes. The nodes are sorted to allow for easy
// lookup in the graph.
func CreateEdge(n1, n2 Node) [2]Node {
	if n1[0] < n2[0] {
		return [2]Node{n1, n2}
	}
	return [2]Node{n2, n1}
}

// HasEdge returns true if the graph has an edge between the two nodes
func (g Graph) HasEdge(n1, n2 Node) bool {
	edge := CreateEdge(n1, n2)
	_, ok := g.Edges[edge]
	return ok
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Parse Input ===============================================
	parse_start := time.Now()
	// secret_numbers := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	// p1 := part1(secret_numbers)
	p1_time := time.Since(p1_start)
	// fmt.Printf("Part 1: %v\n", p1)

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
