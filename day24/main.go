// Most of this code is translated from
// https://github.com/python/cpython/blob/3.13/Lib/graphlib.py
package main

import (
	_ "embed"
	"fmt"
	"time"
)

// ID is the three letters that make up the name of a node.
type ID [3]rune

func (id ID) String() string {
	return fmt.Sprintf("%c%c%c", id[0], id[1], id[2])
}

// node_info stores information about a given node
type node_info struct {
	// id is the name of the given node
	id ID

	// npredecessors
	npredecessors int

	// successors is the slice of successor nodes. Can hold duplicates, as long as
	// they're all reflected in the successor's npredecessors count
	successors []ID
}

func new_node_info(node ID) node_info {
	return node_info{node, 0, make([]ID, 0)}
}

type TopoSorter struct {
	// node2info contains the map of IDs to info
	node2info    map[ID]node_info
	ready_nodes  []ID
	n_passed_out int
	n_finished   int
}

// get_node_info finds the information for an existing node. If the node is not already
// in the map, then it will add it.
// In the python code, this function is used to both add non-existent items to
// ts.node2info, AND alter the values in ts.node2info
func (ts *TopoSorter) get_node_info(node ID) *node_info {
	result, exists := ts.node2info[node]
	if !exists {
		result = new_node_info(node)
		ts.node2info[node] = result
	}
	return &result
}

// Add will add this node to the TopoSorter.
//
// If called multiple times with the same node argument, the set of dependencies will
// be the union of all dependencies pass in.
//
// It is possible to add a node with no dependencies, as well as provide a dependency
// twice. If a node that has not been provided before is included among `predecessors`
// it will be automatically added to the graph with no predecessors of its own.
//
// Throws an error if called after "prepare"
func (ts *TopoSorter) Add(node ID, predecessors ...ID) error {
	if len(ts.ready_nodes) == 0 {
		return fmt.Errorf("Nodes cannot be added after a call to prepare.")
	}

	// Create the node -> predecessor edges
	node_info := ts.get_node_info(node)
	node_info.npredecessors += len(predecessors)

	// Create the predecessor -> node edges
	for _, pred := range predecessors {
		pred_info := ts.get_node_info(pred)
		pred_info.successors = append(pred_info.successors, node)
	}

	return nil
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
