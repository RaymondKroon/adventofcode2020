package main

import (
	"adventofcode2020"
	"fmt"
	"sort"
)

func hasSum(preamble []int, sum int) bool {
	sorted := make([]int, len(preamble))
	copy(sorted, preamble)
	sort.Ints(sorted)
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] >= sum {
			continue
		}
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i]+sorted[j] == sum {
				return true
			}
		}
	}
	return false
}

func SolvePart1(numbers []int, preambleSize int) int {
	for i := preambleSize; i < len(numbers)-preambleSize; i++ {
		preamble := numbers[i-preambleSize : i]
		if !hasSum(preamble, numbers[i]) {
			return numbers[i]
		}
	}
	return -1
}

func SolvePart2(numbers []int, target int) int {
	for sliceSize := 2; sliceSize < len(numbers); sliceSize++ {
		for i := 0; i < len(numbers)-sliceSize; i++ {
			if adventofcode2020.Sum(numbers[i:i+sliceSize]) == target {
				return adventofcode2020.Min(numbers[i:i+sliceSize]) + adventofcode2020.Max(numbers[i:i+sliceSize])
			}
		}
	}
	return -1
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	lines, _ := adventofcode2020.ReadInputLines("./input/day09.txt")
	numbers, _ := adventofcode2020.Atoi(lines)
	p1 := SolvePart1(numbers, 25)
	fmt.Printf("(part1) %d\n", p1)
	p2 := SolvePart2(numbers, p1)
	fmt.Printf("(part2) %d\n", p2)

}
