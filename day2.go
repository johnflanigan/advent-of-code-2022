package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines() []string {
	file, err := os.Open("input2.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()
	check(err)

	return lines
}

func part1() {
	lines := readLines()

	score := 0

	// Opponent: A for Rock, B for Paper, and C for Scissors
	// Player: X for Rock, Y for Paper, and Z for Scissors

	win := make(map[string]string)
	win["X"] = "C"
	win["Y"] = "A"
	win["Z"] = "B"

	tie := make(map[string]string)
	tie["X"] = "A"
	tie["Y"] = "B"
	tie["Z"] = "C"

	value := make(map[string]int)
	value["X"] = 1
	value["Y"] = 2
	value["Z"] = 3

	// 1 for Rock, 2 for Paper, and 3 for Scissors

	for _, line := range lines {
		opponent := string(line[0])
		player := string(line[2])

		if win[player] == opponent {
			score += 6
		} else if tie[player] == opponent {
			score += 3
		}

		score += value[player]
	}

	fmt.Println("Part one:", score)
}

func part2() {
	lines := readLines()

	score := 0

	// Opponent: A for Rock, B for Paper, and C for Scissors
	// Required outcome: X means you need to lose,
	//                   Y means you need to end the round in a draw,
	//                   and Z means you need to win

	move := make(map[string]map[string]string)
	move["A"] = make(map[string]string)
	move["A"]["X"] = "C" // Rock beats scissors
	move["A"]["Y"] = "A" // Rock ties rock
	move["A"]["Z"] = "B" // Rock loses to paper

	move["B"] = make(map[string]string)
	move["B"]["X"] = "A" // Paper beats rock
	move["B"]["Y"] = "B" // Paper ties paper
	move["B"]["Z"] = "C" // Paper loses to scissors

	move["C"] = make(map[string]string)
	move["C"]["X"] = "B" // Scissors beats paper
	move["C"]["Y"] = "C" // Scissors ties scissors
	move["C"]["Z"] = "A" // Scissors loses to rock

	// 1 for Rock, 2 for Paper, and 3 for Scissors
	value := make(map[string]int)
	value["A"] = 1
	value["B"] = 2
	value["C"] = 3
	value["X"] = 0
	value["Y"] = 3
	value["Z"] = 6

	for _, line := range lines {
		opponent := string(line[0])
		outcome := string(line[2])

		player := move[opponent][outcome]

		score += value[player]
		score += value[outcome]
	}

	fmt.Println("Part two:", score)
}
