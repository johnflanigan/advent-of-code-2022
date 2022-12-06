package main

import (
	"fmt"
)

func (d Day) Day6() {
	day6part1()
	day6part2()
}

// Searches line for a window of unique characters.
// Returns the index that the window ends at.
// Panics if no window is found.
func findUniqueChars(line string, windowSize int) int {
	counts := make(map[int]int)

	for i := 0; i < windowSize; i++ {
		val := int(line[i])
		counts[val] = counts[val] + 1
	}

	if len(counts) == windowSize {
		return windowSize - 1
	}

	for i := windowSize; i < len(line); i++ {
		// Remove element now outside window
		prev := int(line[i-windowSize])
		counts[prev] = counts[prev] - 1
		if counts[prev] == 0 {
			delete(counts, prev)
		}

		// Add new element in window
		current := int(line[i])
		counts[current] = counts[current] + 1

		if len(counts) == windowSize {
			return i
		}
	}

	panic(fmt.Sprintf("Could not %d adjacent unique characters", windowSize))
}

func day6part1() {
	lines := readLines("input6.txt")

	line := lines[0]

	index := findUniqueChars(line, 4)

	fmt.Println("Part one:", index+1)
}

func day6part2() {
	lines := readLines("input6.txt")

	line := lines[0]

	index := findUniqueChars(line, 14)

	fmt.Println("Part one:", index+1)
}
