package main

import (
	"adventofcode2020"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	stringInput, _ := adventofcode2020.ReadInput("../input/day01.txt")
	input, _ := adventofcode2020.Atoi(stringInput)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	stringInput, _ := adventofcode2020.ReadInput("../input/day01.txt")
	input, _ := adventofcode2020.Atoi(stringInput)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2sorted(input)
	}
}
