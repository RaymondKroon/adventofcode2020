package main

import (
	"adventofcode2020/util"
	"fmt"
	"regexp"
)

type BusLine struct {
	Id     int
	Offset int
}

func parseBusLines(input string) (result []BusLine) {
	result = make([]BusLine, 0)
	regex := regexp.MustCompile(`([0-9]+|x)(?:,|$)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	for i, m := range matches {
		if m[1] != "x" {
			result = append(result, BusLine{util.MustAtoi(m[1]), i})
		}
	}

	return result
}

func solvePart1(earliest int, busLines []BusLine) int {
	leaveTimes := make(map[int]int, len(busLines))
	for _, bl := range busLines {
		t := ((earliest / bl.Id) + 1) * bl.Id
		leaveTimes[bl.Id] = t
	}

	var min int
	var busId int
	first := true
	for k, v := range leaveTimes {
		if first || v < min {
			min = v
			busId = k
			first = false
		}
	}

	return busId * (min - earliest)

}

func solvePart2(busLines []BusLine) int {
	result := 0
	step := 1

	for _, bl := range busLines {
		for (result+bl.Offset)%bl.Id != 0 {
			result += step
		}

		step *= bl.Id
	}

	return result
}

func main() {
	defer util.Stopwatch("Run")()
	lines, _ := util.ReadInputLines("./input/day13.txt")
	earliestDeparture := util.MustAtoi(lines[0])
	busLines := parseBusLines(lines[1])

	fmt.Println("(part1)", solvePart1(earliestDeparture, busLines)) // 3997
	fmt.Println("(part2)", solvePart2(busLines))                    // 3997
}
