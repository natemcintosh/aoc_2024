package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test_input string = `2333133121414131402`
var small_input string = `12345`

func TestCreateDiskSmall(t *testing.T) {
	want := []int{
		0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2,
	}
	got, _ := create_disk(small_input)
	assert.Equal(t, want, got)
}

func TestCreateDisk(t *testing.T) {
	want := []int{
		0, 0,
		-1, -1, -1,
		1, 1, 1,
		-1, -1, -1,
		2,
		-1, -1, -1,
		3, 3, 3,
		-1,
		4, 4,
		-1,
		5, 5, 5, 5,
		-1,
		6, 6, 6, 6,
		-1,
		7, 7, 7,
		-1,
		8, 8, 8, 8,
		9, 9}
	got, _ := create_disk(test_input)
	assert.Equal(t, want, got)
}

func TestPart1Small(t *testing.T) {
	disk, _ := create_disk(small_input)
	got := part1(disk)
	want := 60
	assert.Equal(t, want, got)
}

func TestPart1(t *testing.T) {
	disk, _ := create_disk(test_input)
	got := part1(disk)
	want := 1928
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	disk, _ := create_disk(raw_text)
	got := part1(disk)
	want := 6446899523367
	assert.Equal(t, want, got)
}

func TestDiskEntryCheckSum(t *testing.T) {
	tests := []struct {
		name string
		de   DiskEntry
		want int
	}{
		{"start at 0", DiskEntry{start: 0, length: 1, file_id: 1}, 0},
		{"start at 0, longer", DiskEntry{start: 0, length: 5, file_id: 1}, 10},

		{"empty", DiskEntry{start: 0, length: 5, file_id: -1}, 0},

		{"empty 2", DiskEntry{start: 2, length: 5, file_id: -1}, 0},

		{"start at 0 2", DiskEntry{start: 0, length: 5, file_id: 2}, 20},
		{"10 to 20", DiskEntry{start: 10, length: 10, file_id: 2}, 290},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.de.CheckSum()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	_, compressed_disk := create_disk(test_input)
	got := part2(compressed_disk)
	want := 2858
	assert.Equal(t, want, got)
}

func TestPart2Real(t *testing.T) {
	_, compressed_disk := create_disk(raw_text)
	got := part2(compressed_disk)
	want := 6478232739671
	assert.Equal(t, want, got)
}
