package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLockKey(t *testing.T) {
	test_cases := []struct {
		input    string
		expected LockKey
		is_lock  bool
	}{
		{`#####
.####
.####
.####
.#.#.
.#...
.....`, LockKey{[5]uint8{0, 5, 3, 4, 3}}, true},
		{
			`#####
##.##
.#.##
...##
...#.
...#.
.....`, LockKey{[5]uint8{1, 2, 0, 5, 3}}, true,
		},
		{`.....
#....
#....
#...#
#.#.#
#.###
#####`, LockKey{[5]uint8{5, 0, 2, 1, 3}}, false},
		{`.....
.....
#.#..
###..
###.#
###.#
#####`, LockKey{[5]uint8{4, 3, 4, 0, 2}}, false},
		{`.....
.....
.....
#....
#.#..
#.#.#
#####`, LockKey{[5]uint8{3, 0, 2, 0, 1}}, false},
	}

	for idx, tc := range test_cases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			lock_key, is_lock := parse_lockey(tc.input)
			assert.Equal(t, tc.expected, lock_key)
			assert.Equal(t, tc.is_lock, is_lock)
		})
	}

}
