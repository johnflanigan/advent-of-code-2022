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
	children  []*directory
}

func (d Day) Day7() {
	day7part1()
	day7part2()
}

func computeSize(root *directory) int {
	root.totalSize += root.files

	for _, child := range root.children {
		root.totalSize += computeSize(child)
	}

	return root.totalSize
}

func buildDirectory(lines []string) directory {
	root := directory{name: "/", totalSize: 0, files: 0, parent: nil, children: []*directory{}}
	current := &root

	for _, line := range lines[1:] {
		//fmt.Println(current)
		if line == "$ cd .." {
			// Add child size to parent size
			//current.parent.size += current.size
			// Move up a directory
			current = current.parent
		} else if line[0:4] == "$ cd" {
			// Create child directory
			child := directory{name: line[5:], totalSize: 0, files: 0, parent: current, children: []*directory{}}
			// Append child to it's parent's children
			current.children = append(current.children, &child)
			// Change to child directory
			current = &child
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

	computeSize(&root)

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
			queue = append(queue, *child)
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
			queue = append(queue, *child)
		}

		if current.totalSize > additionalSpaceRequired && current.totalSize < result {
			result = current.totalSize
		}
	}

	fmt.Println("Part two:", result)
}
