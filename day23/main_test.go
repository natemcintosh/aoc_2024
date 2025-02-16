package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const test_input = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func TestNewNode(t *testing.T) {
	tests := []struct {
		input    string
		want     Node
		want_err error
	}{
		{"ab", Node{'a', 'b'}, nil},
		{"cd", Node{'c', 'd'}, nil},
		{"", Node{}, BadNodeInput},
		{"abc", Node{}, BadNodeInput},
	}
	for _, tc := range tests {
		got, err := NewNode(tc.input)
		if tc.want_err != nil {
			assert.Equal(t, tc.want_err, err)
		}
		assert.Equal(t, tc.want, got)
	}
}

func BenchmarkNewNode(b *testing.B) {
	for b.Loop() {
		NewNode("ab")
	}
}

func TestParse(t *testing.T) {
	input := `ab-cd
zs-cd`
	want := Graph{
		Edges: map[[2]Node]struct{}{
			{Node{'a', 'b'}, Node{'c', 'd'}}: {},
			{Node{'c', 'd'}, Node{'z', 's'}}: {},
		},
		Nodes: map[Node]int{
			{'a', 'b'}: 1,
			{'c', 'd'}: 2,
			{'z', 's'}: 1,
		}}
	got := parse(input)
	assert.Equal(t, want, got)
}

func BenchmarkParse(b *testing.B) {
	input := `ab-cd
zs-cd`
	for b.Loop() {
		parse(input)
	}
}

func BenchmarkHasEdge(b *testing.B) {
	g := parse(test_input)
	for b.Loop() {
		g.HasEdge(Node{'a', 'b'}, Node{'c', 'd'})
	}
}

func TestCompareNodes(t *testing.T) {
	tests := []struct {
		a, b Node
		want int
	}{
		{Node{'a', 'b'}, Node{'a', 'b'}, 0},
		{Node{'b', 'd'}, Node{'b', 'd'}, 0},
		{Node{'a', 'b'}, Node{'b', 'a'}, -1},
		{Node{'a', 'b'}, Node{'a', 'c'}, -1},
		{Node{'a', 'b'}, Node{'c', 'd'}, -1},
		{Node{'a', 'b'}, Node{'b', 'c'}, -1},
		{Node{'a', 'c'}, Node{'a', 'b'}, 1},
		{Node{'b', 'a'}, Node{'a', 'a'}, 1},
	}
	for _, tc := range tests {
		got := compare_nodes(tc.a, tc.b)
		assert.Equal(t, tc.want, got)
	}
}

func TestPart1(t *testing.T) {
	g := parse(test_input)
	want := 7
	got := part1(g)
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	g := parse(raw_text)
	got := part1(g)
	want := 1411
	assert.Equal(t, want, got)
}

func BenchmarkPart1(b *testing.B) {
	benchmarks := []struct {
		name  string
		graph Graph
	}{
		{"test", parse(test_input)},
		{"real", parse(raw_text)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for b.Loop() {
				part1(bm.graph)
			}
		})
	}
}

func TestAddIfPossible(t *testing.T) {
	g := parse(`ka-co
ta-co
de-co
ta-ka
de-ta
ka-de`)
	tests := []struct {
		name string
		fc   FullyConnected
		n    Node
		want bool
	}{
		{"first", FullyConnected{[]Node{{'k', 'a'}}}, Node{'c', 'o'}, true},
		{"second", FullyConnected{[]Node{{'k', 'a'}, {'c', 'o'}}}, Node{'t', 'a'}, true},
		{"third", FullyConnected{[]Node{{'k', 'a'}, {'c', 'o'}}}, Node{'z', 'z'}, false},
		{"fourth", FullyConnected{[]Node{{'k', 'a'}, {'c', 'o'}, {'t', 'a'}}}, Node{'d', 'e'}, true},
	}

	for _, tc := range tests {
		got := tc.fc.AddIfPossible(tc.n, g)
		assert.Equal(t, tc.want, got)

		// Make sure that the node was added to the FCG
		if tc.want {
			assert.Contains(t, tc.fc.Nodes, tc.n)
		}
	}
}

func TestPart2(t *testing.T) {
	g := parse(test_input)
	want := "co,de,ka,ta"
	got := part2(g)
	assert.Equal(t, want, got)
}
