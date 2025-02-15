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
	for b.Loop() {
		step_n(2024, 2000)
	}
}

func BenchmarkPart1(b *testing.B) {
	nums := parse(raw_text)
	for b.Loop() {
		part1(nums)
	}
}

func TestPart1(t *testing.T) {
	want := 37327623
	got := part1([]int{1, 10, 100, 2024})
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	nums := parse(raw_text)
	want := 14622549304
	got := part1(nums)
	assert.Equal(t, want, got)
}

func TestTrackChanges(t *testing.T) {
	secret := 123
	changes2 := make(map[[4]int]int)
	changes := track_all_changes_for_seller(secret, 10, changes2)

	want := map[[4]int]int{
		{-3, 6, -1, -1}: 4, {6, -1, -1, 0}: 4,
		{-1, -1, 0, 2}: 6, {-1, 0, 2, -2}: 4,
		{0, 2, -2, 0}: 4, {2, -2, 0, -2}: 2,
	}
	assert.Equal(t, want, changes)
}

func TestNewChangeTracker(t *testing.T) {
	want := ChangeTracker{[2]int{3, 0}, [4]int{0, 0, 0, -3}}
	got := NewChangeTracker([2]int{3, 0})
	assert.Equal(t, want, got)
}

func TestPart2(t *testing.T) {
	want := 23
	got := part2([]int{1, 2, 3, 2024})
	assert.Equal(t, want, got)
}

func TestPart2Real(t *testing.T) {
	nums := parse(raw_text)
	want := 1735
	got := part2(nums)
	assert.Equal(t, want, got)
}

func BenchmarkPart2(b *testing.B) {
	nums := parse(raw_text)
	for b.Loop() {
		part2(nums)
	}
}
