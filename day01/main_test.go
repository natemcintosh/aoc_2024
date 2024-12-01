package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	raw_input := `3   4
4   3
2   5
1   3
3   9
3   3`
	l, r := parse(raw_input)
	got := part1(l, r)
	want := 11
	assert.Equal(t, want, got)
}

func TestPart1RealInput(t *testing.T) {
	l, r := parse(raw_text)
	got := part1(l, r)
	want := 1646452
	assert.Equal(t, want, got)
}

func TestPart2(t *testing.T) {
	raw_input := `3   4
4   3
2   5
1   3
3   9
3   3`
	l, r := parse(raw_input)
	got := part2(l, r)
	want := 31
	assert.Equal(t, want, got)

	// Again, with part2_v2
	got = part2_v2(l, r)
	assert.Equal(t, want, got)
}

func TestPart2RealInput(t *testing.T) {
	l, r := parse(raw_text)
	got := part2(l, r)
	want := 23609874
	assert.Equal(t, want, got)

	// Again, with part2_v2
	got = part2_v2(l, r)
	assert.Equal(t, want, got)
}

// Set up a benchmark table, for the purpose of comparing the two part2 functions in
// their speed and allocations.
func BenchmarkPart2(b *testing.B) {
	l, r := parse(raw_text)
	benchmarks := []struct {
		name string
		fn   func([]int, []int) int
	}{
		{"part2", part2},
		{"part2_v2", part2_v2},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(l, r)
			}
		})
	}
}
