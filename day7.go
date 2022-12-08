package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type directory struct {
	name      string
	totalSize int
	files     int
	parent    *directory
	children  []directory
}

func (d Day) Day7() {
	day7part1()
	day7part2()
}

func updateSize(root directory) directory {
	root.totalSize += root.files

	for i := range root.children {
		root.children[i] = updateSize(root.children[i])
		root.totalSize += root.children[i].totalSize
	}

	return root
}

func buildDirectory(lines []string) directory {
	root := directory{name: "/", totalSize: 0, files: 0, children: []directory{}}
	current := &root

	for _, line := range lines[1:] {
		if line == "$ cd .." {
			// Move up a directory
			current = current.parent
		} else if line[0:4] == "$ cd" {
			// Create child directory
			child := directory{name: line[5:], parent: current, children: []directory{}}
			// Append child to it's parent's children
			current.children = append(current.children, child)
			// Change to child directory
			current = &current.children[len(current.children)-1]
		} else {
			r, _ := regexp.Compile("(\\d+) \\S+")
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				size, err := strconv.Atoi(matches[1])
				check(err)
				current.files += size
			}
		}
	}

	// We compute the directory sizes at the end rather than when we move up a directory. Otherwise, we miss some
	// directories because the terminal output does not return to root at the end.
	root = updateSize(root)

	return root
}

func day7part1() {
	lines := readLines("input7.txt")
	root := buildDirectory(lines)

	queue := []directory{root}
	total := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, child := range current.children {
			queue = append(queue, child)
		}

		if current.totalSize < 100000 {
			total += current.totalSize
		}
	}

	fmt.Println("Part one:", total)
}

func day7part2() {
	lines := readLines("input7.txt")
	root := buildDirectory(lines)

	totalDiskSpace := 70000000
	usedDiskSpace := root.totalSize
	unusedDiskSpace := totalDiskSpace - usedDiskSpace

	requiredDiskSpace := 30000000
	additionalSpaceRequired := requiredDiskSpace - unusedDiskSpace

	queue := []directory{root}
	result := math.MaxInt

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, child := range current.children {
			queue = append(queue, child)
		}

		if current.totalSize > additionalSpaceRequired && current.totalSize < result {
			result = current.totalSize
		}
	}

	fmt.Println("Part two:", result)
}
