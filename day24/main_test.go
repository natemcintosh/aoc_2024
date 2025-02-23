package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FuzzNameFromID is a fuzz test that validates encoding and decoding of an ID.
// func FuzzNameFromID(f *testing.F) {
// 	// Seed with some valid 3-character ASCII strings.
// 	for r1 := rune(32); r1 < 127; r1++ {
// 		f.Add(r1, r1+1, r1+2)
// 	}

// 	f.Fuzz(func(t *testing.T, r1, r2, r3 rune) {
// 		// Check if the runes are valid ASCII. Also skip early ASCII characters before
// 		// space, as they are not printable.
// 		if r1 > 127 || r2 > 127 || r3 > 127 || r1 < 32 || r2 < 32 || r3 < 32 {
// 			t.Skip("non-ASCII character encountered")
// 		}
// 		id := Wire{r1, r2, r3}
// 		// Encode then decode the ID.
// 		recovered := NameFromInt(id.ID())

// 		// Check if the original and recovered IDs match.
// 		if id != recovered {
// 			t.Errorf("expected %+v, got %+v", id, recovered)
// 		}
// 	})
// }

func BenchmarkID(b *testing.B) {
	// Create an ID, then encode and decode it.
	id := Wire{'a', 'b', 'c'}
	for b.Loop() {
		NameFromInt(id.ID())
	}
}

func TestNewID(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"abc", false},
		{"def", false},
		{"ghi", false},
		{"jklm", true},
		{"", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				assert.Panics(t, func() { NewWire(tt.name) })
				return
			}
		})
	}
}

// TestConstructDiGraph tests the creation of the map
// ┌──A──┐
// ▼     ▼
// B     C
// │     │
// │     ▼
// └────►D
func TestConstructDiGraph(t *testing.T) {
	// Create a new directed g
	g := NewDiGraph()

	// Add the edges
	g.AddEdge(NewWire("AAA"), NewWire("BBB"))
	g.AddEdge(NewWire("AAA"), NewWire("CCC"))
	g.AddEdge(NewWire("BBB"), NewWire("DDD"))
	g.AddEdge(NewWire("CCC"), NewWire("DDD"))

	// Check that the internal maps are as expected
	// First the from-to slices, in alphabetical order: A, B, C, D
	// From A
	assert.Equal(t, []Wire{NewWire("BBB"), NewWire("CCC")}, g.from_to[NewWire("AAA")])

	// From B
	assert.Equal(t, []Wire{NewWire("DDD")}, g.from_to[NewWire("BBB")])

	// From C
	assert.Equal(t, []Wire{NewWire("DDD")}, g.from_to[NewWire("CCC")])

	// From D
	assert.Equal(t, []Wire{}, g.from_to[NewWire("DDD")])

	// To A
	assert.Equal(t, []Wire{}, g.to_from[NewWire("AAA")])

	// To B
	assert.Equal(t, []Wire{NewWire("AAA")}, g.to_from[NewWire("BBB")])

	// To C
	assert.Equal(t, []Wire{NewWire("AAA")}, g.to_from[NewWire("CCC")])

	// To D
	assert.Equal(t, []Wire{NewWire("BBB"), NewWire("CCC")}, g.to_from[NewWire("DDD")])
}

// TestTopoSort tests that we get the proper ordering for this map
// ┌──A──┐
// ▼     ▼
// B     C
// │     │
// │     ▼
// └────►D
func TestTopoSort(t *testing.T) {
	a := NewWire("AAA")
	b := NewWire("BBB")
	c := NewWire("CCC")
	d := NewWire("DDD")

	// Create and populate the graph
	g := NewDiGraph()
	g.AddEdge(a, b)
	g.AddEdge(a, c)
	g.AddEdge(b, d)
	g.AddEdge(c, d)

	want := []Wire{a, c, b, d}
	got, err := g.TopoSort()
	if err != nil {
		t.Fatal("Found a cycle when should not have")
	}
	assert.Equal(t, want, got)
}
