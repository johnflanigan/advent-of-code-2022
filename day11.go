package main

import (
	"fmt"
	"sort"
)

type monkey struct {
	operation func(int) int
	test      func(int) int
}

func (d Day) Day11() {
	day11part1()
	day11part2()
}

// TODO read in input from file instead of hardcoding input
func day11part1() {
	items := inputStartingItems()
	monkeys := inputMonkeys()

	inspections := make([]int, len(items))

	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := j % len(monkeys)

			holding := items[monkey]
			items[monkey] = []int{}

			for _, item := range holding {
				inspections[monkey]++
				newLevel := monkeys[monkey].operation(item) / 3
				nextMonkey := monkeys[monkey].test(newLevel)
				items[nextMonkey] = append(items[nextMonkey], newLevel)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Println("Part one:", inspections[0]*inspections[1])
}

func day11part2() {
	items := inputStartingItems()
	monkeys := inputMonkeys()

	inspections := make([]int, len(items))

	lcd := 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19
	//lcd := 13 * 17 * 19 * 23

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := j % len(monkeys)

			holding := items[monkey]
			items[monkey] = []int{}

			for _, item := range holding {
				inspections[monkey]++
				newLevel := monkeys[monkey].operation(item)

				if newLevel > lcd {
					newLevel = newLevel % lcd
				}

				nextMonkey := monkeys[monkey].test(newLevel)

				items[nextMonkey] = append(items[nextMonkey], newLevel)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	fmt.Println("Part two:", inspections[0]*inspections[1])
}

func testStartingItems() map[int][]int {
	items := make(map[int][]int)
	items[0] = []int{79, 98}
	items[1] = []int{54, 65, 75, 74}
	items[2] = []int{79, 60, 97}
	items[3] = []int{74}
	return items
}

func inputStartingItems() map[int][]int {
	items := make(map[int][]int)
	items[0] = []int{72, 97}
	items[1] = []int{55, 70, 90, 74, 95}
	items[2] = []int{74, 97, 66, 57}
	items[3] = []int{86, 54, 53}
	items[4] = []int{50, 65, 78, 50, 62, 99}
	items[5] = []int{90}
	items[6] = []int{88, 92, 63, 94, 96, 82, 53, 53}
	items[7] = []int{70, 60, 71, 69, 77, 70, 98}
	return items
}

func testMonkeys() map[int]monkey {
	monkeys := make(map[int]monkey)

	monkeys[0] = monkey{
		operation: func(old int) int {
			return old * 19
		},
		test: func(item int) int {
			if item%23 == 0 {
				return 2
			} else {
				return 3
			}
		},
	}
	monkeys[1] = monkey{
		operation: func(old int) int {
			return old + 6
		},
		test: func(item int) int {
			if item%19 == 0 {
				return 2
			} else {
				return 0
			}
		},
	}
	monkeys[2] = monkey{
		operation: func(old int) int {
			return old * old
		},
		test: func(item int) int {
			if item%13 == 0 {
				return 1
			} else {
				return 3
			}
		},
	}
	monkeys[3] = monkey{
		operation: func(old int) int {
			return old + 3
		},
		test: func(item int) int {
			if item%17 == 0 {
				return 0
			} else {
				return 1
			}
		},
	}
	return monkeys
}

func inputMonkeys() map[int]monkey {
	monkeys := make(map[int]monkey)

	monkeys[0] = monkey{
		operation: func(old int) int {
			return old * 13
		},
		test: func(item int) int {
			if item%19 == 0 {
				return 5
			} else {
				return 6
			}
		},
	}
	monkeys[1] = monkey{
		operation: func(old int) int {
			return old * old
		},
		test: func(item int) int {
			if item%7 == 0 {
				return 5
			} else {
				return 0
			}
		},
	}
	monkeys[2] = monkey{
		operation: func(old int) int {
			return old + 6
		},
		test: func(item int) int {
			if item%17 == 0 {
				return 1
			} else {
				return 0
			}
		},
	}
	monkeys[3] = monkey{
		operation: func(old int) int {
			return old + 2
		},
		test: func(item int) int {
			if item%13 == 0 {
				return 1
			} else {
				return 2
			}
		},
	}
	monkeys[4] = monkey{
		operation: func(old int) int {
			return old + 3
		},
		test: func(item int) int {
			if item%11 == 0 {
				return 3
			} else {
				return 7
			}
		},
	}
	monkeys[5] = monkey{
		operation: func(old int) int {
			return old + 4
		},
		test: func(item int) int {
			if item%2 == 0 {
				return 4
			} else {
				return 6
			}
		},
	}
	monkeys[6] = monkey{
		operation: func(old int) int {
			return old + 8
		},
		test: func(item int) int {
			if item%5 == 0 {
				return 4
			} else {
				return 7
			}
		},
	}
	monkeys[7] = monkey{
		operation: func(old int) int {
			return old * 7
		},
		test: func(item int) int {
			if item%3 == 0 {
				return 2
			} else {
				return 3
			}
		},
	}
	return monkeys
}
