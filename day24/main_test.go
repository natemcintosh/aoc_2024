package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Real(t *testing.T) {
	got := part1()
	want := 36035961805936
	assert.Equal(t, want, got)
}
