package main

import (
	"fmt"
	"math"
)

func (d Day) Day8() {
	day8part1()
	day8part2()
}

func getVisibleFromTop(heights [][]int) [][]bool {
	rows := len(heights)
	cols := len(heights[0])

	visible := make([][]bool, rows)
	for row := 0; row < rows; row++ {
		visible[row] = make([]bool, cols)
	}

	for col := 0; col < cols; col++ {
		max := math.MinInt

		for row := 0; row < rows; row++ {
			if heights[row][col] > max {
				visible[row][col] = true
				max = heights[row][col]
			}
		}
	}

	return visible
}

func getVisibleFromBottom(heights [][]int) [][]bool {
	rows := len(heights)
	cols := len(heights[0])

	visible := make([][]bool, rows)
	for row := 0; row < rows; row++ {
		visible[row] = make([]bool, cols)
	}

	for col := 0; col < cols; col++ {
		max := math.MinInt

		for row := rows - 1; row >= 0; row-- {
			if heights[row][col] > max {
				visible[row][col] = true
				max = heights[row][col]
			}
		}
	}

	return visible
}

func getVisibleFromLeft(heights [][]int) [][]bool {
	rows := len(heights)
	cols := len(heights[0])

	visible := make([][]bool, rows)
	for row := 0; row < rows; row++ {
		visible[row] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		max := math.MinInt

		for col := 0; col < cols; col++ {
			if heights[row][col] > max {
				visible[row][col] = true
				max = heights[row][col]
			}
		}
	}

	return visible
}

func getVisibleFromRight(heights [][]int) [][]bool {
	rows := len(heights)
	cols := len(heights[0])

	visible := make([][]bool, rows)
	for row := 0; row < rows; row++ {
		visible[row] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		max := math.MinInt

		for col := cols - 1; col >= 0; col-- {
			if heights[row][col] > max {
				visible[row][col] = true
				max = heights[row][col]
			}
		}
	}

	return visible
}

func getVisible(heights [][]int) [][]bool {
	top := getVisibleFromTop(heights)
	bottom := getVisibleFromBottom(heights)
	left := getVisibleFromLeft(heights)
	right := getVisibleFromRight(heights)

	rows := len(heights)
	cols := len(heights[0])

	visible := make([][]bool, rows)
	for row := 0; row < rows; row++ {
		visible[row] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			visible[row][col] = top[row][col] || bottom[row][col] || left[row][col] || right[row][col]
		}
	}

	return visible
}

func getHeights(lines []string) [][]int {
	rows := len(lines)
	cols := len(lines[0])

	heights := make([][]int, rows)
	for row := 0; row < rows; row++ {
		heights[row] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			heights[row][col] = int(lines[row][col]) - '0'
		}
	}
	return heights
}

func day8part1() {
	lines := readLines("input8.txt")
	heights := getHeights(lines)
	visible := getVisible(heights)

	count := 0
	for _, row := range visible {
		for _, el := range row {
			if el {
				count++
			}
		}
	}

	fmt.Println("Part one:", count)
}

func getUpScore(heights [][]int, row int, col int) int {
	// Walk up counting how many trees are less than this tree. When we encounter a tree greater than or equal, we must
	// add that tree to the score, then return.
	current := heights[row][col]
	score := 0
	for i := row - 1; i >= 0; i-- {
		if heights[i][col] >= current {
			score++
			break
		}
		score++
	}

	return score
}

func getDownScore(heights [][]int, row int, col int) int {
	rows := len(heights)

	// Walk down counting how many trees are less than this tree.
	current := heights[row][col]
	score := 0
	for i := row + 1; i < rows; i++ {
		if heights[i][col] >= current {
			score++
			break
		}
		score++
	}

	return score
}

func getLeftScore(heights [][]int, row int, col int) int {
	// Walk left counting how many trees are less than this tree.
	current := heights[row][col]
	score := 0
	for j := col - 1; j >= 0; j-- {
		if heights[row][j] >= current {
			score++
			break
		}
		score++
	}

	return score
}

func getRightScore(heights [][]int, row int, col int) int {
	cols := len(heights[0])

	// Walk right counting how many trees are less than this tree.
	current := heights[row][col]
	score := 0
	for j := col + 1; j < cols; j++ {
		if heights[row][j] >= current {
			score++
			break
		}
		score++
	}

	return score
}

func getScores(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	scores := make([][]int, rows)
	for row := 0; row < rows; row++ {
		scores[row] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			up := getUpScore(heights, row, col)
			down := getDownScore(heights, row, col)
			left := getLeftScore(heights, row, col)
			right := getRightScore(heights, row, col)

			scores[row][col] = up * down * left * right
		}
	}

	return scores
}

func day8part2() {
	lines := readLines("input8.txt")
	heights := getHeights(lines)
	scores := getScores(heights)

	max := math.MinInt
	for _, row := range scores {
		for _, el := range row {
			if el > max {
				max = el
			}
		}
	}

	fmt.Println("Part two:", max)
}
