package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/natemcintosh/aoc_2024/utils"
)

var pattern *regexp.Regexp = regexp.MustCompile(`(\d+) ([b|r|g])`)

// parse_game takes a game. For example, the record of a few games might look like this:
//
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
//
// In game 1, three sets of cubes are revealed from the bag (and then put back again).
// The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green
// cubes, and 6 blue cubes; the third set is only 2 green cubes.
func parse_game(game string) []RGB {
	// Split into each draw from the bag
	draws := strings.Split(game, ";")

	// Break it into easy to digest final chunks
	pairings := make([]RGB, len(draws))
	for grab_idx, grab := range draws {
		// Get the matches
		matches := pattern.FindAllStringSubmatch(grab, -1)
		pairings[grab_idx] = parse_rgb(matches)

	}

	return pairings
}

func parse_rgb(matches [][]string) RGB {
	rgb := RGB{}

	// Iterate over the matches, parsing them, and editing the `rgb`
	for _, m := range matches {
		num, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatalf("Could not parse %v", m[1])
		}
		switch m[2] {
		case "r":
			rgb.R = num
		case "g":
			rgb.G = num
		case "b":
			rgb.B = num
		default:
			log.Fatalln("Letter was not r, g, or b")
		}
	}

	return rgb
}

// RGB represents the number of colored balls seen
type RGB struct {
	R, G, B int
}

func (self RGB) Contains(other RGB) bool {
	return (self.R >= other.R) && (self.G >= other.G) && (self.B >= other.B)
}

func (self RGB) Power() int {
	return self.R * self.G * self.B
}

// The Elf would first like to know which games would have been possible if the bag
// contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
//
// In the example above, games 1, 2, and 5 would have been possible if the bag had been
// loaded with that configuration. However, game 3 would have been impossible because
// at one point the Elf showed you 20 red cubes at once; similarly, game 4 would also
// have been impossible because the Elf showed you 15 blue cubes at once. If you add up
// the IDs of the games that would have been possible, you get 8.
func part1(games [][]RGB) int {
	var result int = 0

	// The max game, over which you should not go
	max_game := RGB{R: 12, G: 13, B: 14}

	// Iterate over the games, checking if they are possible
	// and if so, adding their index to the result
	for game_idx, game := range games {
		// A flag for marking if a game is bad
		is_bad := false
		// Iterate over the draws, checking if they are possible
		// and if not, break out of this loop
		for _, draw := range game {
			if !max_game.Contains(draw) {
				is_bad = true
				break
			}
		}

		// After checking all the cases, if the game is not bad, add it to the result
		if !is_bad {
			result += game_idx + 1
		}
	}

	return result
}

// The Elf says they've stopped producing snow because
// they aren't getting any water! He isn't sure why the
// water stopped; however, he can show you how to get to
// the water source to check it out for yourself. It's just up ahead!
//
// As you continue your walk, the Elf poses a second question:
// in each game you played, what is the fewest number of cubes
// of each color that could have been in the bag to make the
// game possible?
func part2(games [][]RGB) int {
	game_powers := 0
	for _, game := range games {
		smallest_rgb := RGB{R: 0, G: 0, B: 0}
		// Get the smallest number of each color
		for _, draw := range game {
			if draw.R > smallest_rgb.R {
				smallest_rgb.R = draw.R
			}
			if draw.G > smallest_rgb.G {
				smallest_rgb.G = draw.G
			}
			if draw.B > smallest_rgb.B {
				smallest_rgb.B = draw.B
			}
		}

		// Add the power of the game
		game_powers += smallest_rgb.Power()
	}

	return game_powers
}

func main() {
	// Time how long it takes to read the file
	// and parse the games
	parse_start := time.Now()

	// === Parse ====================================================
	raw_text := utils.ReadFile("example_day/input.txt")
	lines := strings.Split(raw_text, "\n")
	games := make([][]RGB, len(lines))
	for i, line := range lines {
		games[i] = parse_game(line)
	}
	p1 := part1(games)
	parse_time := time.Now().Sub(parse_start)

	// === Part 1 ====================================================
	p1_start := time.Now()
	p1_time := time.Now().Sub(p1_start)
	fmt.Printf("Part 1: %v\n", p1)

	// === Part 2 ====================================================
	p2_start := time.Now()
	p2 := part2(games)
	p2_time := time.Now().Sub(p2_start)
	fmt.Printf("Part 2: %v\n", p2)

	// === Print Results ============================================
	fmt.Printf("\n\nSetup took %v\n", parse_time)
	fmt.Printf("Part 1 took %v\n", p1_time)
	fmt.Printf("Part 2 took %v\n", p2_time)

}
