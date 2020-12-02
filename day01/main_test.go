package main

import (
	"adventofcode2020"
	"testing"
)

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]int) int
	}{
		{name: "part1", fn: part1},
		{name: "part2", fn: part2},
		{name: "part2sorted", fn: part2sorted},
	}

	stringInput, _ := adventofcode2020.ReadInput("../input/day01.txt")
	input, _ := adventofcode2020.Atoi(stringInput)

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(input)
			}
		})
	}
}
