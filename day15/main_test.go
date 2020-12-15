package main

import (
	"fmt"
	"testing"
)

func TestPlayGame(t *testing.T) {
	type args struct {
		start []int
		turns int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{start: []int{0, 3, 6}, turns: 2020}, 436},
		{args{start: []int{1, 3, 2}, turns: 2020}, 1},
		{args{start: []int{2, 1, 3}, turns: 2020}, 10},
		{args{start: []int{1, 2, 3}, turns: 2020}, 27},
		{args{start: []int{2, 3, 1}, turns: 2020}, 78},
		{args{start: []int{3, 2, 1}, turns: 2020}, 438},
		{args{start: []int{3, 1, 2}, turns: 2020}, 1836},
		{args{start: []int{0, 3, 6}, turns: 30000000}, 175594},
		{args{start: []int{1, 3, 2}, turns: 30000000}, 2578},
		{args{start: []int{2, 1, 3}, turns: 30000000}, 3544142},
		{args{start: []int{1, 2, 3}, turns: 30000000}, 261214},
		{args{start: []int{2, 3, 1}, turns: 30000000}, 6895259},
		{args{start: []int{3, 2, 1}, turns: 30000000}, 18},
		{args{start: []int{3, 1, 2}, turns: 30000000}, 362},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%d", tt.args.start, tt.args.turns), func(t *testing.T) {
			if got := PlayGame(tt.args.start, tt.args.turns); got != tt.want {
				t.Errorf("PlayGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PlayGame([]int{3, 1, 2}, 30000000)
	}
}
