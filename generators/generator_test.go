package main

import (
	"reflect"
	"testing"

	. "github.com/dave/jennifer/jen"
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
		got, err := ParseGate(tc.line)
		if err != nil {
			t.Fatalf("parse_gate(%q) returned error: %v", tc.line, err)
		}
		if reflect.DeepEqual(got, tc.want) == false {
			t.Errorf("parse_gate(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}
