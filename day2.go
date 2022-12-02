package main

import (
	"fmt"
)

func (d Day) Day2() {
	day2part1()
	day2part2()
}

func mod(a, b int) int {
	return (a%b + b) % b
}

// TODO eliminate the need for adding dayX to each method
func day2part1() {
	lines := readLines("input2.txt")

	score := 0

	// Opponent: A for Rock, B for Paper, C for Scissors
	// Player: X for Rock, Y for Paper, Z for Scissors

	for _, line := range lines {
		// Convert inputs such that 0 for Rock, 1 for Paper, 2 for Scissors
		opponent := int(line[0]) - 65
		player := int(line[2]) - 88

		result := mod(opponent-player, 3)

		if result == 2 {
			score += 6
		} else if result == 0 {
			score += 3
		}

		// Add bonus points (1 point for Rock, 2 points for Paper, 3 points for Scissors)
		score += player + 1
	}

	fmt.Println("Part one:", score)
}

func day2part2() {
	lines := readLines("input2.txt")

	score := 0

	// Opponent: A for Rock, B for Paper, C for Scissors
	// Result: X means you need to lose, Y means you need to tie, Z means you need to win

	for _, line := range lines {
		// Convert input such that 0 for Rock, 1 for Paper, 2 for Scissors
		opponent := int(line[0]) - 65
		// Convert input such that -1 for Lose, 0 for Tie, 1 for Win
		result := int(line[2]) - 89

		player := mod(opponent+result, 3)

		if result == 1 {
			score += 6
		} else if result == 0 {
			score += 3
		}

		score += player + 1
	}

	fmt.Println("Part two:", score)
}
