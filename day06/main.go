package main

import (
	"adventofcode2020/util"
	"fmt"
	"strings"
)

type Counters struct {
	Any int
	All int
}

func CountAnswers(input string) Counters {
	answers := make(map[string]int)
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		for _, c := range l {
			answers[string(c)] += 1
		}
	}

	all := 0
	for _, v := range answers {
		if v == len(lines) {
			all += 1
		}
	}

	return Counters{
		Any: len(answers),
		All: all,
	}
}

func main() {
	defer util.Stopwatch("Run")()
	text, _ := util.ReadInput("./input/day06.txt")
	groupChunks := strings.Split(text, "\n\n")
	any := 0
	all := 0
	for _, g := range groupChunks {
		counts := CountAnswers(g)
		any += counts.Any
		all += counts.All
	}
	fmt.Printf("(part1) total: %d\n", any)
	fmt.Printf("(part2) total: %d\n", all)
}
