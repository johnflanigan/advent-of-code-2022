package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Day struct{}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(name string) []string {
	file, err := os.Open(name)
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

func main() {
	arg := os.Args[1]
	day, err := strconv.Atoi(arg)
	// TODO provide more helpful information on usage
	check(err)

	fmt.Println("Running day:", day)

	d := Day{}
	name := fmt.Sprintf("Day%d", day)
	meth := reflect.ValueOf(d).MethodByName(name)
	meth.Call(nil)
}
