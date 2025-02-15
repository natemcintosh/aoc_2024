package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	for i := 0; i < b.N; i++ {
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
