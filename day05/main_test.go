package main

import (
	"adventofcode2020/util"
	"testing"
)

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(s string) BoardingPass
	}{
		{name: "calculate", fn: ParseBoardingPassOld},
		{name: "bitshift", fn: ParseBoardingPass},
	}
	stringInput, _ := util.ReadInputLines("../input/day05.txt")
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, input := range stringInput {
					bm.fn(input)
				}
			}
		})
	}
}
