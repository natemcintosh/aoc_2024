package main

import (
	"testing"

	. "github.com/dave/jennifer/jen"
	"github.com/natemcintosh/aoc_2024/circuits"
	"github.com/stretchr/testify/assert"
)

func TestParse_gate(t *testing.T) {
	test_cases := []struct {
		line string
		want Code
	}{
		{
			"y33 AND x33 -> bfn",
			Id("bfn").Op(":=").Id("y33").Op("&&").Id("x33"),
		},
		{
			"y32 XOR x32 -> rck",
			Id("rck").Op(":=").Id("y32").Op("!=").Id("x32"),
		},
		{
			"x30 OR y30 -> gns",
			Id("gns").Op(":=").Id("x30").Op("||").Id("y30"),
		},
	}

	for _, tc := range test_cases {
		got, _, err := ParseGate(tc.line)
		if err != nil {
			t.Fatalf("parse_gate(%q) returned error: %v", tc.line, err)
		}
		assert.Equal(t, tc.want, got)
	}
}

func TestParseInputAssignment(t *testing.T) {
	test_cases := []struct {
		line string
		want Code
	}{
		{"x00: 1", Id("x00").Op(":=").Id("x").Index(Lit(0))},
		{"x01: 0", Id("x01").Op(":=").Id("x").Index(Lit(1))},
		{"y00: 1", Id("y00").Op(":=").Id("y").Index(Lit(0))},
		{"y01: 1", Id("y01").Op(":=").Id("y").Index(Lit(1))},
		{"y13: 0", Id("y13").Op(":=").Id("y").Index(Lit(13))},
		{"a13: 0", Id("a13").Op(":=").Id("a").Index(Lit(13))},
	}

	for _, tc := range test_cases {
		got, err := ParseInputAssignment(tc.line)
		if err != nil {
			t.Fatalf("parse_input_assignment(%q) returned error: %v", tc.line, err)
		}
		assert.Equal(t, tc.want, got)
	}
}

func BenchmarkCircuit(b *testing.B) {
	// Get the input values
	x, y := circuits.InputValues()

	// Benchmark the circuit
	for b.Loop() {
		circuits.Circuit(x, y)
	}
}
