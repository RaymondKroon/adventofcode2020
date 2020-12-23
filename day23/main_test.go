package main

import (
	"testing"
)

func TestCrabCups_Move(t *testing.T) {
	start := CrabCups{3, 8, 9, 1, 2, 5, 4, 6, 7}
	m1 := CrabCups{2, 8, 9, 1, 5, 4, 6, 7, 3}
	m2 := CrabCups{5, 4, 6, 7, 8, 9, 1, 3, 2}
	m3 := CrabCups{8, 9, 1, 3, 4, 6, 7, 2, 5}
	m4 := CrabCups{4, 6, 7, 9, 1, 3, 2, 5, 8}
	m5 := CrabCups{1, 3, 6, 7, 9, 2, 5, 8, 4}
	m6 := CrabCups{9, 3, 6, 7, 2, 5, 8, 4, 1}
	m7 := CrabCups{2, 5, 8, 3, 6, 7, 4, 1, 9}
	m8 := CrabCups{6, 7, 4, 1, 5, 8, 3, 9, 2}
	m9 := CrabCups{5, 7, 4, 1, 8, 3, 9, 2, 6}
	m10 := CrabCups{8, 3, 7, 4, 1, 9, 2, 6, 5}

	for i, move := range [][]CrabCups{
		{start, m1},
		{m1, m2},
		{m2, m3},
		{m3, m4},
		{m4, m5},
		{m5, m6},
		{m6, m7},
		{m7, m8},
		{m8, m9},
		{m9, m10},
	} {
		result := move[0].Move()
		if result != move[1] {
			t.Errorf("move error %d: %v to %v, result == %v", i+1, move[0], move[1], result)
		}
	}
}
