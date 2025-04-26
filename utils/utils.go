package utils

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// ReadFile reads a file and returns its content as a string
func ReadFile(filename string) string {
	byte_contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(byte_contents))
}

// ParseInt converts a string to an int and panics if it fails
func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// ParseFloat converts a string to a float64 and panics if it fails
func ParseFloat(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// ParseBool converts a string to a bool and panics if it fails
func ParseBool(s string) bool {
	s = strings.TrimSpace(s)
	if s == "1" || strings.ToLower(s) == "true" {
		return true
	}
	if s == "0" || strings.ToLower(s) == "false" {
		return false
	}
	panic("invalid boolean string")
}

// GetGroups is essentially the same as FindAllStringSubmatch, but it returns
// just the groups and not the full match
func GetGroups(re *regexp.Regexp, s string) [][]string {
	matches := re.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		match = slices.Delete(match, 0, 1)
	}
	return matches
}

// DiGraph is a directed graph that stores the edges in both directions. This is
// necessary to be able to traverse the graph in both directions. The graph is
// represented as a map of edges.
type DiGraph[T comparable] struct {
	from_to map[T][]T
	to_from map[T][]T
}

func NewDiGraph[T comparable]() DiGraph[T] {
	return DiGraph[T]{
		from_to: make(map[T][]T),
		to_from: make(map[T][]T),
	}
}

// AddEdge adds an edge to the graph. It adds the edge in both directions.
func (g *DiGraph[T]) AddEdge(from, to T) {
	g.from_to[from] = append(g.from_to[from], to)
	g.to_from[to] = append(g.to_from[to], from)

	// Make sure that there is an entry (even if it is empty) for `to` in the from_to
	// map. This is necessary to be able to traverse the graph in both directions.
	if _, ok := g.from_to[to]; !ok {
		g.from_to[to] = []T{}
	}

	// Make sure that there is an entry (even if it is empty) for `from` in the to_from
	// map. This is necessary to be able to traverse the graph in both directions.
	if _, ok := g.to_from[from]; !ok {
		g.to_from[from] = []T{}
	}
}

// RemoveEdge removes an edge from the graph. It removes the edge in both directions.
func (g *DiGraph[T]) RemoveEdge(from, to T) {
	g.from_to[from] = slices.DeleteFunc(g.from_to[from], func(w T) bool { return w == to })
	g.to_from[to] = slices.DeleteFunc(g.to_from[to], func(w T) bool { return w == from })
}

// HasEdges checks if there are any edges in the DiGraph at all.
func (g *DiGraph[T]) HasEdges() bool {
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

// TopoSort performs a topological sort on the graph. It returns a slice of items in
// topological order. It returns an error if the graph is not a DAG (Directed Acyclic
// Graph). This is performed using [Kahn's Algorithm](wikipedia.org/Topological_sorting#Kahn's_algorithm)
// Note that the order may vary from one call to the next, but is guaranteed to be
// correctly sorted.
func (g *DiGraph[T]) TopoSort() ([]T, error) {
	// Create an empty slices that wil container all the sorted elements
	L := make([]T, 0, len(g.from_to))

	// Create a slices that we will iterate over. Start with all nodes that have no
	// incoming edges
	S := make([]T, 0)

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
		return []T{}, fmt.Errorf("no nodes with no incoming edges")
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
		return []T{}, fmt.Errorf("There is a cycle in the graph.")
	}

	return L, nil
}
