package main

import (
	"reflect"
	"strings"
	"testing"
)

func tile(input string) Tile {
	return ParseTile(strings.Split(input, "\n"))
}

var (
	original = tile(
		`#.#.......
..........
..........
..........
..........
..........
..........
..........
..........
#.........`)

	flipped = tile(
		`.......#.#
..........
..........
..........
..........
..........
..........
..........
..........
.........#`)

	rotated = tile(
		`#........#
..........
.........#
..........
..........
..........
..........
..........
..........
..........`)

	symmetric = tile(
		`#.#.......
..........
..........
..........
..........
..........
..........
..........
..........
.......#.#`)
)

func TestTile_Rotate(t1 *testing.T) {
	type fields struct {
		pixels PixelsType
	}
	tests := []struct {
		name   string
		fields fields
		want   *Tile
	}{
		{name: "rotate right", fields: fields{pixels: original.pixels}, want: &rotated},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tile{
				pixels: tt.fields.pixels,
			}
			if got := t.RotateCW(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("RotateCW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_Flip(t1 *testing.T) {
	type fields struct {
		pixels PixelsType
	}
	tests := []struct {
		name   string
		fields fields
		want   *Tile
	}{
		{name: "flip Horizontal", fields: fields{pixels: original.pixels}, want: &flipped},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tile{
				pixels: tt.fields.pixels,
			}
			if got := t.FlipH(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("FlipH() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_GetAllPositions(t *testing.T) {
	tests := []struct {
		name string
		tile Tile
		want int
	}{
		{name: "non-symmetric", tile: original, want: 8},
		{name: "symmetric", tile: symmetric, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			positions := original.GetAllPositions()
			unique := make(map[PixelsType]bool)
			for _, p := range positions {
				unique[p.pixels] = true
			}

			if len(unique) != 8 {
				t1.Errorf("All_Positions = %d, should be %d", len(unique), tt.want)
			}
		})
	}
}

func TestTile_Neighbours(t *testing.T) {
	tests := []struct {
		name   string
		first  Tile
		second Tile
		top    bool
		right  bool
		bottom bool
		left   bool
	}{
		{"self", original, original, false, false, false, false},
		{"symmetric", symmetric, symmetric, false, false, false, false},
		{"flipped", original, flipped, false, true, false, true},
		{"flipped symmetric", symmetric, *symmetric.FlipH(), true, true, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			if tt.first.IsTopNeighbour(&tt.second) != tt.top {
				t1.Errorf("%s, should be %t", "top", tt.top)
			}
			if tt.first.IsRightNeighbour(&tt.second) != tt.right {
				t1.Errorf("%s, should be %t", "right", tt.right)
			}
			if tt.first.IsBottomNeighbour(&tt.second) != tt.bottom {
				t1.Errorf("%s, should be %t", "bottom", tt.bottom)
			}
			if tt.first.IsLeftNeighbour(&tt.second) != tt.left {
				t1.Errorf("%s, should be %t", "left", tt.left)
			}
		})
	}
}
