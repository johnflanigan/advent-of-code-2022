package main

import (
	"fmt"
	"strconv"
)

type point struct {
	row int
	col int
}

var exists = struct{}{}

func (d Day) Day9() {
	day9part1()
	day9part2()
}

func day9part1() {
	lines := readLines("input9.txt")

	// starting position
	rope := [2]point{}
	for i := 0; i < len(rope); i++ {
		rope[i] = point{row: 0, col: 0}
	}

	rowCol := make(map[int]map[int]struct{})

	for _, line := range lines {
		fmt.Println(line)
		direction := line[0]
		count, err := strconv.Atoi(line[2:])
		check(err)

		for i := 0; i < count; i++ {
			// Update head
			rope[0] = updateHead(direction, rope[0])

			// Update remainder of rope
			for j := 1; j < len(rope); j++ {
				rope[j] = updateKnot(rope[j-1], rope[j])
			}

			// Record tail
			tail := rope[len(rope)-1]
			set, prs := rowCol[tail.row]
			if !prs {
				set = make(map[int]struct{})
			}
			set[tail.col] = exists
			rowCol[tail.row] = set
		}
	}

	count := 0
	for _, value := range rowCol {
		count += len(value)
	}

	fmt.Println("Part one:", count)
}

func updateKnot(prev point, current point) point {
	rowDiff := prev.row - current.row
	colDiff := prev.col - current.col
	rowDist := abs(rowDiff)
	colDist := abs(colDiff)

	if (rowDist > 1 && colDist > 0) || (rowDist > 0 && colDist > 1) {
		if rowDiff > 0 {
			current.row += 1
		} else {
			current.row -= 1
		}
		if colDiff > 0 {
			current.col += 1
		} else {
			current.col -= 1
		}
	} else if rowDist > 1 {
		if rowDiff > 0 {
			current.row += 1
		} else {
			current.row -= 1
		}
	} else if colDist > 1 {
		if colDiff > 0 {
			current.col += 1
		} else {
			current.col -= 1
		}
	}

	return current
}

func updateHead(direction uint8, head point) point {
	if direction == 'U' {
		head.row += 1
	} else if direction == 'D' {
		head.row -= 1
	} else if direction == 'R' {
		head.col += 1
	} else if direction == 'L' {
		head.col -= 1
	}

	return head
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func move(direction uint8, head point, tail point) (point, point) {
	if direction == 'U' {
		head.row += 1
	} else if direction == 'D' {
		head.row -= 1
	} else if direction == 'R' {
		head.col += 1
	} else if direction == 'L' {
		head.col -= 1
	}

	rowDiff := head.row - tail.row
	colDiff := head.col - tail.col
	rowDist := abs(rowDiff)
	colDist := abs(colDiff)

	if (rowDist > 1 && colDist > 0) || (rowDist > 0 && colDist > 1) {
		if rowDiff > 0 {
			tail.row += 1
		} else {
			tail.row -= 1
		}
		if colDiff > 0 {
			tail.col += 1
		} else {
			tail.col -= 1
		}
	} else if rowDist > 1 {
		if rowDiff > 0 {
			tail.row += 1
		} else {
			tail.row -= 1
		}
	} else if colDist > 1 {
		if colDiff > 0 {
			tail.col += 1
		} else {
			tail.col -= 1
		}
	}

	// Check
	rowDiff = head.row - tail.row
	colDiff = head.col - tail.col
	rowDist = abs(rowDiff)
	colDist = abs(colDiff)
	if (rowDist > 1 && colDist > 0) || (rowDist > 0 && colDist > 1) || (rowDist > 1) || (colDist > 1) {
		panic("This should never be true")
	}

	return head, tail
}

// TODO refactor to reuse code
func day9part2() {
	lines := readLines("input9.txt")

	// starting position
	rope := [10]point{}
	for i := 0; i < len(rope); i++ {
		rope[i] = point{row: 0, col: 0}
	}

	rowCol := make(map[int]map[int]struct{})

	for _, line := range lines {
		fmt.Println(line)
		direction := line[0]
		count, err := strconv.Atoi(line[2:])
		check(err)

		for i := 0; i < count; i++ {
			// Update head
			rope[0] = updateHead(direction, rope[0])

			// Update remainder of rope
			for j := 1; j < len(rope); j++ {
				rope[j] = updateKnot(rope[j-1], rope[j])
			}

			// Record tail
			tail := rope[len(rope)-1]
			set, prs := rowCol[tail.row]
			if !prs {
				set = make(map[int]struct{})
			}
			set[tail.col] = exists
			rowCol[tail.row] = set
		}
	}

	count := 0
	for _, value := range rowCol {
		count += len(value)
	}

	fmt.Println("Part two:", count)
}
