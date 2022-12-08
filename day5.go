package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type step struct {
	count  int
	source int
	target int
}

// TODO Fix program structure so rows in day8 does not shadow rows in day5. This would also enable us to avoid having
// to append the day to each part1() and part2() function.
const day5rows = 8

func (d Day) Day5() {
	day5part1()
	day5part2()
}

func getStacks(lines []string) map[int][]int {
	stacks := make(map[int][]int)

	for index := 1; index < len(lines[0]); index += 4 {
		row := (index / 4) + 1
		stacks[row] = []int{}
	}

	// 3 represents the number of day5rows of input for the stack

	for i := 0; i < day5rows; i++ {
		line := lines[i]

		for index := 1; index < len(line); index += 4 {
			if line[index] >= 'A' && line[index] <= 'Z' {
				row := (index / 4) + 1
				stacks[row] = append(stacks[row], int(line[index]))
			}
		}
	}

	return stacks
}

func getStep(line string) step {
	r, _ := regexp.Compile("move (\\d+) from (\\d+) to (\\d+)")
	matches := r.FindStringSubmatch(line)

	count, err := strconv.Atoi(matches[1])
	check(err)
	source, err := strconv.Atoi(matches[2])
	check(err)
	target, err := strconv.Atoi(matches[3])
	check(err)

	return step{count: count, source: source, target: target}
}

func rearrangeOneAtOnce(lines []string, stacks map[int][]int) {
	for i := day5rows + 2; i < len(lines); i++ {
		step := getStep(lines[i])

		for j := 0; j < step.count; j++ {
			top := stacks[step.source][0]
			stacks[step.source] = stacks[step.source][1:]
			stacks[step.target] = append([]int{top}, stacks[step.target]...)
		}
	}
}

func rearrangeManyAtOnce(lines []string, stacks map[int][]int) {
	for i := day5rows + 2; i < len(lines); i++ {
		step := getStep(lines[i])

		top := stacks[step.source][0:step.count]
		stacks[step.source] = stacks[step.source][step.count:]

		// Must make a copy of slice before appending to it. Otherwise, we are inadvertently appending to the middle of
		// stacks[step.source]
		result := make([]int, len(top))
		copy(result, top)
		stacks[step.target] = append(result, stacks[step.target]...)
	}
}

func day5part1() {
	lines := readLines("input5.txt")
	stacks := getStacks(lines)

	rearrangeOneAtOnce(lines, stacks)

	result := ""
	for index := 1; index < len(lines[0]); index += 4 {
		row := (index / 4) + 1
		result += string(stacks[row][0])
	}

	fmt.Println("Part one:", result)
}

func day5part2() {
	lines := readLines("input5.txt")
	stacks := getStacks(lines)

	rearrangeManyAtOnce(lines, stacks)

	result := ""
	for index := 1; index < len(lines[0]); index += 4 {
		row := (index / 4) + 1
		result += string(stacks[row][0])
	}

	fmt.Println("Part two:", result)
}
