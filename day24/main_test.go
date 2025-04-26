package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FuzzNameFromID is a fuzz test that validates encoding and decoding of an ID.
func FuzzNameFromID(f *testing.F) {
	// Seed with some valid 3-character ASCII strings.
	for r1 := rune(32); r1 < 127; r1++ {
		f.Add(r1, r1+1, r1+2)
	}

	f.Fuzz(func(t *testing.T, r1, r2, r3 rune) {
		// Check if the runes are valid ASCII. Also skip early ASCII characters before
		// space, as they are not printable.
		if r1 > 127 || r2 > 127 || r3 > 127 || r1 < 32 || r2 < 32 || r3 < 32 {
			t.Skip("non-ASCII character encountered")
		}
		id := Wire{r1, r2, r3}
		// Encode then decode the ID.
		recovered := NameFromInt(id.ID())

		// Check if the original and recovered IDs match.
		if id != recovered {
			t.Errorf("expected %+v, got %+v", id, recovered)
		}
	})
}

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
