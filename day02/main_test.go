package main

import (
	"adventofcode2020"
	"testing"
)

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]PasswordLine) int
	}{{name: "part1", fn: Part1CountCorrectPasswords},
		{name: "part2", fn: Part2CountCorrectPasswords},
	}

	stringInput, _ := adventofcode2020.ReadInputLines("../input/day02.txt")
	input, _ := ParseInput(stringInput)

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(input)
			}
		})
	}
}
