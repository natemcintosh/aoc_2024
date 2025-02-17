package main

import (
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"
)

// Node represents a node in a graph. It's made up of two runes
type Node [2]rune

var BadNodeInput = fmt.Errorf("Could not produce a Node\n")

func (n Node) Format(f fmt.State, c rune) {
	switch c {
	case 'v': // default format
		fmt.Fprintf(f, "%c%c", n[0], n[1])
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
		// fmt.Printf("Tried to parse into Node but failed: %v\n", s)
		return Node{}, BadNodeInput
	}
	var n Node
	for idx, char := range s {
		n[idx] = char
	}

	return n, nil
}

// Graph is an undirected graph. It stores both the nodes, and the edges between nodes.
// Note that the Nodes map also stores the number of neighbors that node has.
type Graph struct {
	Edges map[[2]Node]struct{}
	Nodes map[Node]int
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
	g := Graph{make(map[[2]Node]struct{}), make(map[Node]int)}
	// For each line, for each pair of letters, add an edge to the graph
	lines := strings.Split(strings.TrimSpace(raw_text), "\n")
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
		g.Nodes[n1] += 1
		g.Nodes[n2] += 1

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

// compare_nodes compares two nodes. Compare by comparing the first rune, then the
// second. A rune with a lower value is considered "less" than a rune with a higher value.
func compare_nodes(a, b Node) int {
	if a[0] < b[0] {
		return -1
	}
	if a[0] > b[0] {
		return 1
	}
	if a[1] < b[1] {
		return -1
	}
	if a[1] > b[1] {
		return 1
	}
	return 0
}

// compare_edges compares two edges. Compare by comparing the first node, then the
// second. A node with a lower value is considered "less" than a node with a higher value.
func compare_edges(a, b [2]Node) int {
	if compare_nodes(a[0], b[0]) == 0 {
		return compare_nodes(a[1], b[1])
	}
	return compare_nodes(a[0], b[0])
}

func part1(g Graph) int {
	// Get a sorted slice of all the edges for easy, deterministic iteration
	sorted_edges := slices.SortedFunc(maps.Keys(g.Edges), compare_edges)

	triplets := make(map[[3]Node]struct{})

	// For each edge, go forward in the list, and check for a common edge between the two nodes
	for idx, e1 := range sorted_edges {
		// If neither node starts with `t`, then continue
		if e1[0][0] != 't' && e1[1][0] != 't' {
			continue
		}

		for _, e2 := range sorted_edges[idx:] {
			// Check if either node in e1 matches either node in e2
			var e3 [2]Node
			if e1[0] == e2[0] {
				e3 = CreateEdge(e1[1], e2[1])
			} else if e1[0] == e2[1] {
				e3 = CreateEdge(e1[1], e2[0])
			} else if e1[1] == e2[0] {
				e3 = CreateEdge(e1[0], e2[1])
			} else if e1[1] == e2[1] {
				e3 = CreateEdge(e1[0], e2[0])
			} else {
				continue
			}

			// If so, check if the graph has a node made up of the other two nodes not
			// yet checked
			if !(g.HasEdge(e3[0], e3[1])) {
				continue
			}

			// If so, sort the nodes, and put them in triplets
			nodes := []Node{e1[0], e1[1], e2[0], e2[1], e3[0], e3[1]}
			slices.SortStableFunc(nodes, compare_nodes)
			nodes = slices.Compact(nodes)
			triplets[[3]Node{nodes[0], nodes[1], nodes[2]}] = struct{}{}
		}
	}

	return len(triplets)
}

// FullyConnected is a set of nodes that are all connected to each other.
type FullyConnected struct {
	Nodes []Node
}

// AddIfPossible will see if `n` is connected to all the nodes in `g.Nodes`. If `n` is
// connected to all the nodes, it will return true, else false.
func (fc *FullyConnected) AddIfPossible(n Node, g Graph) bool {
	for _, gn := range fc.Nodes {
		if !g.HasEdge(gn, n) {
			return false
		}
	}
	return true
}

func part2(g Graph) string {
	// Create a pool of fully connected sub-graphs
	fcg_pool := make([]FullyConnected, 0, len(g.Nodes))

	// Iterate over the nodes
	for n := range g.Nodes {
		// Attempt to add it to all existing FCGs
		for idx, fcg := range fcg_pool {
			can_add := fcg.AddIfPossible(n, g)
			if can_add {
				fcg_pool[idx].Nodes = append(fcg_pool[idx].Nodes, n)
			}
		}

		// Add a new FCG that has just this one.
		fcg_pool = append(fcg_pool, FullyConnected{[]Node{n}})
	}

	// Sort the fully connected graphs by how many nodes they have, most first
	slices.SortStableFunc(fcg_pool, func(a, b FullyConnected) int {
		if len(a.Nodes) > len(b.Nodes) {
			return -1
		} else if len(a.Nodes) < len(b.Nodes) {
			return 1
		}
		return 0
	})

	// Get the nodes in the largest FCG
	biggest_fcg := fcg_pool[0].Nodes
	slices.SortStableFunc(biggest_fcg, compare_nodes)

	// Get all the nodes, and join them
	var res strings.Builder
	for idx, n := range biggest_fcg {
		res.WriteRune(n[0])
		res.WriteRune(n[1])
		if idx < len(biggest_fcg)-1 {
			res.WriteRune(',')
		}
	}

	return res.String()
}

// The input text of the puzzle
//
//go:embed input.txt
var raw_text string

func main() {
	// === Parse Input ===============================================
	parse_start := time.Now()
	graph := parse(raw_text)
	parse_time := time.Since(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1 := part1(graph)
	p1_time := time.Since(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(graph)
	p2_time := time.Since(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)
}
