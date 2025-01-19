package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMix(t *testing.T) {
	want := 37
	got := mix(42, 15)
	assert.Equal(t, want, got)
}

func TestPrune(t *testing.T) {
	want := 16113920
	got := prune(100000000)
	assert.Equal(t, want, got)
}

func TestStepSequential(t *testing.T) {
	secret := 123
	tests := []struct {
		want int
	}{
		{15887950}, {16495136}, {527345}, {704524},
		{1553684}, {12683156}, {11100544}, {12249484},
		{7753432}, {5908254},
	}

	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.want), func(t *testing.T) {
			got := step(secret)
			assert.Equal(t, tc.want, got)
			secret = got
		})
	}
}

func TestStepN(t *testing.T) {
	tests := []struct {
		name   string
		secret int
		n      int
		want   int
	}{
		{"1", 1, 2000, 8685429},
		{"10", 10, 2000, 4700978},
		{"100", 100, 2000, 15273692},
		{"2024", 2024, 2000, 8667524},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := step_n(tc.secret, tc.n)
			assert.Equal(t, tc.want, got)
		})
	}
}

func BenchmarkStepN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		step_n(2024, 2000)
	}
}

func TestPart1(t *testing.T) {
	want := 37327623
	got := part1([]int{1, 10, 100, 2024})
	assert.Equal(t, want, got)
}
