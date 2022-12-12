package main

import (
	"fmt"
	"math"
)

func (d Day) Day12() {
	day12part1()
	day12part2()
}

func getIndex(row int, col int, cols int) int {
	return row*cols + col
}

func getRowCol(index int, cols int) (int, int) {
	row := index / cols
	col := index % cols
	return row, col
}

func getNeighbors(index int, rows int, cols int) []int {
	row, col := getRowCol(index, cols)

	var neighbors []int

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := col + direction[1]
		if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols {
			neighbors = append(neighbors, getIndex(nextRow, nextCol, cols))
		}
	}

	return neighbors
}

func day12part1() {
	lines := readLines("input12.txt")

	rows := len(lines)
	cols := len(lines[0])

	// Read input into slice
	vertices := make(map[int]int)
	for row, line := range lines {
		for col, char := range line {
			index := getIndex(row, col, cols)
			vertices[index] = int(char)
		}
	}

	start := 0
	end := 0
	for key, value := range vertices {
		if value == 'S' {
			start = key
		} else if value == 'E' {
			end = key
		}
	}
	vertices[start] = 'a'
	vertices[end] = 'z'

	dist := make(map[int]int)
	prev := make(map[int]int)
	q := make(map[int]struct{})
	// for each vertex v in Graph.Vertices:
	for i := 0; i < len(vertices); i++ {
		// dist[v] ← INFINITY
		dist[i] = math.MaxInt
		// prev[v] ← UNDEFINED
		// add v to Q
		q[i] = exists
	}

	// dist[source] ← 0
	dist[start] = 0

	// while Q is not empty:
	for len(q) > 0 {
		// u ← vertex in Q with min dist[u]
		u := findMin(q, dist)
		// remove u from Q
		delete(q, u)

		// for each neighbor v of u still in Q:
		for _, v := range getNeighbors(u, rows, cols) {
			_, prs := q[v]
			if prs && vertices[u] >= vertices[v]-1 {
				// alt ← dist[u] + Graph.Edges(u, v)
				alt := dist[u] + 1
				// if alt < dist[v]:
				if alt < dist[v] {
					// dist[v] ← alt
					dist[v] = alt
					// prev[v] ← u
					prev[v] = u
				}
			}
		}
	}

	fmt.Println("Part one:", dist[end])
}

func findMin(q map[int]struct{}, dist map[int]int) int {
	minIndex := 0
	minDist := math.MaxInt
	for u := range q {
		if dist[u] <= minDist {
			minIndex = u
			minDist = dist[u]
		}
	}
	return minIndex
}

func day12part2() {
	lines := readLines("input12.txt")

	rows := len(lines)
	cols := len(lines[0])

	// Read input into slice
	vertices := make(map[int]int)
	for row, line := range lines {
		for col, char := range line {
			index := getIndex(row, col, cols)
			vertices[index] = int(char)
		}
	}

	start := []int{}
	end := 0
	for key, value := range vertices {
		if value == 'S' || value == 'a' {
			start = append(start, key)
		} else if value == 'E' {
			end = key
		}
	}

	dist := make(map[int]int)
	prev := make(map[int]int)
	q := make(map[int]struct{})
	// for each vertex v in Graph.Vertices:
	for i := 0; i < len(vertices); i++ {
		// dist[v] ← INFINITY
		dist[i] = math.MaxInt
		// prev[v] ← UNDEFINED
		// add v to Q
		q[i] = exists
	}

	for _, key := range start {
		// dist[source] ← 0
		dist[key] = 0
		vertices[key] = 'a'
	}
	vertices[end] = 'z'

	// while Q is not empty:
	for len(q) > 0 {
		// u ← vertex in Q with min dist[u]
		u := findMin(q, dist)
		// remove u from Q
		delete(q, u)

		// for each neighbor v of u still in Q:
		for _, v := range getNeighbors(u, rows, cols) {
			_, prs := q[v]
			if prs && vertices[u] >= vertices[v]-1 {
				// alt ← dist[u] + Graph.Edges(u, v)
				alt := dist[u] + 1
				// if alt < dist[v]:
				if alt < dist[v] {
					// dist[v] ← alt
					dist[v] = alt
					// prev[v] ← u
					prev[v] = u
				}
			}
		}
	}

	fmt.Println("Part two:", dist[end])
}
