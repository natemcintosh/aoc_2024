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
		Nodes: map[Node]struct{}{
			{'a', 'b'}: {},
			{'c', 'd'}: {},
			{'z', 's'}: {},
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
