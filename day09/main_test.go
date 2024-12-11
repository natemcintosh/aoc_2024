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
	got := create_disk(small_input)
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
	got := create_disk(test_input)
	assert.Equal(t, want, got)
}

func TestPart1Small(t *testing.T) {
	disk := create_disk(small_input)
	got := part1(disk)
	want := 60
	assert.Equal(t, want, got)
}

func TestPart1(t *testing.T) {
	disk := create_disk(test_input)
	got := part1(disk)
	want := 1928
	assert.Equal(t, want, got)
}

func TestPart1Real(t *testing.T) {
	disk := create_disk(raw_text)
	got := part1(disk)
	want := 6446899523367
	assert.Equal(t, want, got)
}
