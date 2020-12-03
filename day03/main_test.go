package main

import (
	"adventofcode2020"
	"testing"
)

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(m Map) int
	}{
		{name: "part1", fn: Part1},
		{name: "part2", fn: Part2},
	}

	for _, bm := range benchmarks {
		stringInput, _ := adventofcode2020.ReadInput("../input/day03.txt")
		m := CreateMap(stringInput)

		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(m)
			}
		})
	}
}
