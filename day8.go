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
	var heights [][]int

	for i, line := range lines {
		heights = append(heights, []int{})
		for _, char := range line {
			heights[i] = append(heights[i], int(char)-'0')
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

func getScoresFromTop(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	scores := make([][]int, rows)
	for row := 0; row < rows; row++ {
		scores[row] = make([]int, cols)
	}

	for col := 0; col < cols; col++ {
		for row := 1; row < rows; row++ {
			// Walk towards the top counting how many trees are less than this tree. When we encounter a tree greater
			// than or equal, we must add that tree to the visible ones, then break.
			current := heights[row][col]
			score := 0
			for i := row - 1; i >= 0; i-- {
				if heights[i][col] >= current {
					score++
					break
				}
				score++
			}

			scores[row][col] = score
		}
	}

	return scores
}

func getScoresFromBottom(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	scores := make([][]int, rows)
	for row := 0; row < rows; row++ {
		scores[row] = make([]int, cols)
	}

	for col := 0; col < cols; col++ {
		for row := rows - 2; row >= 0; row-- {
			// Walk towards the bottom counting how many trees are less than this tree.
			current := heights[row][col]
			score := 0
			for i := row + 1; i < rows; i++ {
				if heights[i][col] >= current {
					score++
					break
				}
				score++
			}
			scores[row][col] = score
		}
	}

	return scores
}

func getScoresFromLeft(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	scores := make([][]int, rows)
	for row := 0; row < rows; row++ {
		scores[row] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 1; col < cols; col++ {
			// Walk towards the left counting how many trees are less than this tree.
			current := heights[row][col]
			score := 0
			for j := col - 1; j >= 0; j-- {
				if heights[row][j] >= current {
					score++
					break
				}
				score++
			}
			scores[row][col] = score
		}
	}

	return scores
}

func getScoresFromRight(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	scores := make([][]int, rows)
	for row := 0; row < rows; row++ {
		scores[row] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := cols - 2; col >= 0; col-- {
			// Walk towards the right counting how many trees are less than this tree.
			current := heights[row][col]
			score := 0
			for j := col + 1; j < cols; j++ {
				if heights[row][j] >= current {
					score++
					break
				}
				score++
			}
			scores[row][col] = score
		}
	}

	return scores
}

func getScores(heights [][]int) [][]int {
	top := getScoresFromTop(heights)
	bottom := getScoresFromBottom(heights)
	left := getScoresFromLeft(heights)
	right := getScoresFromRight(heights)

	rows := len(heights)
	cols := len(heights[0])

	scores := make([][]int, rows)
	for row := 0; row < rows; row++ {
		scores[row] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			scores[row][col] = top[row][col] * bottom[row][col] * left[row][col] * right[row][col]
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
