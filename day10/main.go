package main

import (
	"adventofcode2020/util"
	"fmt"
	"sort"
)

func Part1(input []int) int {
	voltages := make([]int, len(input))
	copy(voltages, input)
	sort.Ints(voltages)
	voltages = append([]int{0}, voltages...)
	voltages = append(voltages, voltages[len(voltages)-1]+3)

	diffs := make(map[int]int)
	for i := 0; i < len(voltages)-1; i++ {
		diffs[voltages[i+1]-voltages[i]] += 1
	}

	return diffs[1] * diffs[3]
}

func Part2(input []int) int {
	voltages := make([]int, len(input))
	copy(voltages, input)
	sort.Ints(voltages)
	voltages = append([]int{0}, voltages...)
	nPaths := make(map[int]int)
	nPaths[0] = 1
	for i := 0; i < len(voltages); i++ {
		v := voltages[i]
		next := voltages[i:]
		for _, step := range []int{1, 2, 3} {
			if util.Contains(next, v) {
				nPaths[v+step] += nPaths[v]
			}
		}
	}

	return nPaths[voltages[len(voltages)-1]]
}

func main() {
	defer util.Stopwatch("Run")()
	lines, _ := util.ReadInputLines("./input/day10.txt")
	voltages, _ := util.StringsAtoi(lines)

	p1 := Part1(voltages)
	fmt.Println("(part1)", p1)

	p2 := Part2(voltages)
	fmt.Println("(part2)", p2)
}
