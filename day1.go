package main

import (
	"fmt"
	"sort"
	"strconv"
)

func (d Day) Day1() {
	day1part1()
	day1part2()
}

func day1part1() {
	lines := readLines("input1.txt")

	max := 0
	sum := 0

	for _, line := range lines {

		if len(line) == 0 {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			i, err := strconv.Atoi(line)
			check(err)

			sum += i
		}
	}

	fmt.Println("Part one:", max)
}

func day1part2() {
	lines := readLines("input1.txt")

	// Initialize array to contain three largest values
	topThree := []int{0, 0, 0}
	sum := 0

	for _, line := range lines {
		if len(line) == 0 {
			// Check if sum is larger than the smallest top three value
			if sum > topThree[0] {
				// Replace the smallest top three value with new sum
				topThree[0] = sum
				// Sort array so index 0 contains the smallest value
				sort.Ints(topThree)
			}
			sum = 0
		} else {
			i, err := strconv.Atoi(line)
			check(err)

			sum += i
		}
	}

	total := 0
	for i := 0; i < 3; i++ {
		total += topThree[i]
	}

	fmt.Println("Part two:", total)
}
