package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	max := 0
	sum := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		text := scanner.Text()

		i, _ := strconv.Atoi(text)

		if len(text) == 0 {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			sum += i
		}
	}

	fmt.Println(max)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part2() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	topThree := []int{0, 0, 0}
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()

		i, _ := strconv.Atoi(text)

		if len(text) == 0 {
			if sum > topThree[0] {
				topThree[0] = sum
				sort.Ints(topThree)
			}
			sum = 0
		} else {
			sum += i
		}
	}

	total := 0
	for i := 0; i < 3; i++ {
		total += topThree[i]
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
