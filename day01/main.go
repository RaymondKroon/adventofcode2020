package main

import (
	"adventofcode2020"
	"sort"
)

func main() {

	stringInput, _ := adventofcode2020.ReadInput("./input/day01.txt")
	input, _ := adventofcode2020.Atoi(stringInput)

	println(part1(input))
	println(part2sorted(input))
}

func part1(input []int) int {
	for iIdx, i := range input {
		for _, j := range input[iIdx+1:] {
			if i+j == 2020 {
				return i * j
			}
		}
	}

	return -1
}

func part2(input []int) int {
	for iIdx, i := range input {
		for jIdx, j := range input[iIdx+1:] {
			if i+j < 2020 {
				for _, k := range input[iIdx+jIdx+1:] {
					if i+j+k == 2020 {
						return i * j * k
					}
				}
			}
		}
	}
	return -1
}

func part2sorted(input []int) int {
	sort.Ints(input)
	for iIdx, i := range input {
		for jIdx, j := range input[iIdx+1:] {
			remainder := 2020 - i - j
			if remainder > 0 {
				elems := input[iIdx+jIdx+1:]
				searchIdx := sort.SearchInts(elems, remainder)
				if searchIdx < len(elems) && elems[searchIdx] == remainder {
					return i * j * remainder
				}
			}
		}
	}
	return -1
}
