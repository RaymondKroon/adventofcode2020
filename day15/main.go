package main

import (
	"adventofcode2020"
	"fmt"
	"regexp"
)

func parseStart(input string) []int {
	regex := regexp.MustCompile(`([0-9]+)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	result := make([]int, len(matches))
	for i := 0; i < len(matches); i++ {
		result[i] = adventofcode2020.MustAtoi(matches[i][1])
	}

	return result
}

func Part1PlayGame(start []int, turns int) int {
	mem := make(map[int]int)
	last := 0
	for i, s := range start {
		last = s
		mem[s] = i + 1
	}
	for i := len(start); i < turns; i++ {
		if lastTs, ok := mem[last]; ok {
			mem[last] = i
			last = i - lastTs
		} else {
			mem[last] = i
			last = 0
		}
	}

	return last
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	input, _ := adventofcode2020.ReadInput("./input/day15.txt")
	start := parseStart(input)
	fmt.Println("(part1)", Part1PlayGame(start, 2020))
	fmt.Println("(part2)", Part1PlayGame(start, 30000000))
}
