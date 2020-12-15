package main

import (
	"adventofcode2020"
	"fmt"
	"regexp"
)

type Int = uint32

func parseStart(input string) []int {
	regex := regexp.MustCompile(`([0-9]+)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	result := make([]int, len(matches))
	for i := 0; i < len(matches); i++ {
		result[i] = adventofcode2020.MustAtoi(matches[i][1])
	}

	return result
}

func PlayGame(start []int, turns int) int {
	mem := make([]Int, turns)
	last := Int(0)
	for i, s := range start {
		last = Int(s)
		mem[s] = Int(i + 1)
	}
	for i := len(start); i < turns; i++ {
		if lastTs := mem[last]; lastTs > 0 {
			mem[last] = Int(i)
			last = Int(i) - lastTs
		} else {
			mem[last] = Int(i)
			last = 0
		}
	}

	return int(last)
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	input, _ := adventofcode2020.ReadInput("./input/day15.txt")
	start := parseStart(input)
	fmt.Println("(part1)", PlayGame(start, 2020))
	fmt.Println("(part2)", PlayGame(start, 30000000))
}
