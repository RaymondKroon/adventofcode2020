package main

import (
	"adventofcode2020"
	"testing"
)

func BenchmarkParser(b *testing.B) {
	lines, _ := adventofcode2020.ReadInputLines("../input/day11.txt")
	for i := 0; i < b.N; i++ {
		parseFloorplan(lines)
	}
}

func BenchmarkSolvers(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(fp FloorPlan) int
	}{
		{name: "part1", fn: solvePart1},
		{name: "part2", fn: solvePart2},
	}

	lines, _ := adventofcode2020.ReadInputLines("../input/day11.txt")
	floorplan := parseFloorplan(lines)

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(floorplan)
			}
		})
	}
}
