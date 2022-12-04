package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type assignment struct {
	start int
	stop  int
}

func (d Day) Day4() {
	day4part1()
	day4part2()
}

func getAssignments(line string) (assignment, assignment) {
	r, _ := regexp.Compile("(\\d+)-(\\d+),(\\d+)-(\\d+)")
	matches := r.FindStringSubmatch(line)

	firstStart, err := strconv.Atoi(matches[1])
	check(err)
	firstStop, err := strconv.Atoi(matches[2])
	check(err)
	secondStart, err := strconv.Atoi(matches[3])
	check(err)
	secondStop, err := strconv.Atoi(matches[4])
	check(err)

	return assignment{start: firstStart, stop: firstStop}, assignment{start: secondStart, stop: secondStop}
}

func day4part1() {
	lines := readLines("input4.txt")

	count := 0

	for _, line := range lines {
		first, second := getAssignments(line)

		if (first.start >= second.start && first.stop <= second.stop) || // first contained within second
			(second.start >= first.start && second.stop <= first.stop) { // second contained within first
			count++
		}
	}

	fmt.Println("Part one:", count)
}

func day4part2() {
	lines := readLines("input4.txt")

	count := 0

	for _, line := range lines {
		first, second := getAssignments(line)

		if (second.start <= first.start && first.start <= second.stop) || // first.start falls within second
			(second.start <= first.stop && first.stop <= second.stop) || // first.stop falls within second
			(first.start <= second.start && second.start <= first.stop) || // second.start falls within first
			(first.start <= second.stop && second.stop <= first.stop) { // second.stop falls within first
			count++
		}
	}

	fmt.Println("Part two:", count)
}
