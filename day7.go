package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// TODO figure out the preferred way of representing this data. Should parent be a *directory or directory? Should
//  children be a []directory or a []*directory.
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

func computeSize(root *directory) int {
	root.totalSize += root.files

	// As someone new to go, this was a tricky bug. I missed that := would create a copy of each element of
	// root.children. Thus, when we passed that reference into compute size, we were no longer referencing the element
	// from root.children, we were referencing the newly created directory. This was a subtle bug that took required
	// stepping through the debugger. The incorrect code is included as a reminder to myself.
	// for _, child := range root.children {
	//     root.totalSize += computeSize(&child)
	// }
	for i := range root.children {
		root.totalSize += computeSize(&root.children[i])
	}

	return root.totalSize
}

// TODO can this be done without pointers?
func buildDirectory(lines []string) *directory {
	root := directory{name: "/", totalSize: 0, files: 0, children: []directory{}}
	current := &root

	for _, line := range lines[1:] {
		if line == "$ cd .." {
			// Move up a directory
			current = current.parent
		} else if line[0:4] == "$ cd" {
			// Create child directory
			child := directory{name: line[5:], totalSize: 0, files: 0, parent: current, children: []directory{}}
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
	computeSize(&root)

	return &root
}

func day7part1() {
	lines := readLines("test.txt")
	root := buildDirectory(lines)

	queue := []*directory{root}
	total := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// This was the same bug encountered in compute size. Because we were using := and range to iterate over
		// current.children, we were creating a new copy of each element of current.children and appending the address
		// to the queue. However, because the child variable was being reused, we were just appending the same memory
		// address to the queue repeatedly.
		// for _, child := range current.children {
		//     queue = append(queue, *child)
		// }
		for i := range current.children {
			queue = append(queue, &current.children[i])
		}

		if current.totalSize < 100000 {
			total += current.totalSize
		}
	}

	fmt.Println("Part one:", total)
}

func day7part2() {
	lines := readLines("test.txt")
	root := buildDirectory(lines)

	totalDiskSpace := 70000000
	usedDiskSpace := root.totalSize
	unusedDiskSpace := totalDiskSpace - usedDiskSpace

	requiredDiskSpace := 30000000
	additionalSpaceRequired := requiredDiskSpace - unusedDiskSpace

	queue := []*directory{root}
	result := math.MaxInt

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := range current.children {
			queue = append(queue, &current.children[i])
		}

		if current.totalSize > additionalSpaceRequired && current.totalSize < result {
			result = current.totalSize
		}
	}

	fmt.Println("Part two:", result)
}
