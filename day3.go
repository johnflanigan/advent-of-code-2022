package main

import (
	"fmt"
)

func (d Day) Day3() {
	day3part1()
	day3part2()
}

func getPriority(char byte) int {
	val := int(char)

	if val >= 65 && val <= 90 {
		return val - 38
	} else if val >= 97 && val <= 122 {
		return val - 96
	}

	panic(fmt.Sprintf("Unexpected character: %c", char))
}

func day3part1() {
	lines := readLines("input3.txt")

	priorities := 0

	for _, line := range lines {
		length := len(line)
		first := [53]bool{}
		second := [53]bool{}

		for i := 0; i < length/2; i++ {
			priority := getPriority(line[i])
			first[priority] = true
		}

		for i := length / 2; i < length; i++ {
			priority := getPriority(line[i])
			second[priority] = true
		}

		for i := range first {
			if first[i] && second[i] {
				priorities += i
			}
		}
	}

	fmt.Println("Part one:", priorities)
}

func day3part2() {
	lines := readLines("input3.txt")

	priorities := 0

	for i := 0; i < len(lines); i += 3 {
		first := [53]bool{}
		second := [53]bool{}
		third := [53]bool{}

		for _, char := range lines[i] {
			priority := getPriority(byte(char))
			first[priority] = true
		}

		for _, char := range lines[i+1] {
			priority := getPriority(byte(char))
			second[priority] = true
		}

		for _, char := range lines[i+2] {
			priority := getPriority(byte(char))
			third[priority] = true
		}

		for i := range first {
			if first[i] && second[i] && third[i] {
				priorities += i
			}
		}
	}

	fmt.Println("Part two:", priorities)
}
