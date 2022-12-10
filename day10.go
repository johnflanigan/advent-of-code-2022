package main

import (
	"fmt"
	"strconv"
)

func (d Day) Day10() {
	day10part1()
	day10part2()
}

func day10part1() {
	lines := readLines("input10.txt")

	signalStrength := 0
	cycle := 0
	sum := 1
	for _, line := range lines {
		if line == "noop" {
			cycle += 1
			if (cycle+20)%40 == 0 {
				signalStrength += cycle * sum
			}
		} else {
			value, err := strconv.Atoi(line[5:])
			check(err)

			cycle += 1
			if (cycle+20)%40 == 0 {
				signalStrength += cycle * sum
			}

			cycle += 1
			if (cycle+20)%40 == 0 {
				signalStrength += cycle * sum
			}
			sum += value
		}
	}

	//fmt.Println(lines)
	//count := 0

	fmt.Println("Part one:", signalStrength)
}

/*
Part two output: BGKAEREZ
###...##..#..#..##..####.###..####.####.
#..#.#..#.#.#..#..#.#....#..#.#.......#.
###..#....##...#..#.###..#..#.###....#..
#..#.#.##.#.#..####.#....###..#.....#...
#..#.#..#.#.#..#..#.#....#.#..#....#....
###...###.#..#.#..#.####.#..#.####.####.
*/
func day10part2() {
	lines := readLines("input10.txt")

	cycle := 0
	pos := 1
	var output []int

	for _, line := range lines {
		if line == "noop" {

			output = drawPixel(cycle, pos, output)
			cycle = (cycle + 1) % 40
		} else {
			value, err := strconv.Atoi(line[5:])
			check(err)

			output = drawPixel(cycle, pos, output)
			cycle = (cycle + 1) % 40

			output = drawPixel(cycle, pos, output)
			cycle = (cycle + 1) % 40

			pos += value
		}
	}

	// Output screen
	fmt.Println("Part two:")
	for i := 0; i < len(output); i++ {
		fmt.Printf("%c", output[i])
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}

func drawPixel(cycle int, pos int, output []int) []int {
	if cycle >= pos-1 && cycle <= pos+1 {
		output = append(output, '#')
	} else {
		output = append(output, '.')
	}
	return output
}
