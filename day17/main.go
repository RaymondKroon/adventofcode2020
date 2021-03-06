package main

import (
	"adventofcode2020/util"
	"fmt"
	"strings"
)

type Int = int32

type Point struct {
	X Int
	Y Int
	Z Int
	W Int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y, p.Z + other.Z, p.W + other.W}
}

func (p Point) Neighbours3d() (points []Point) {
	delta := []Int{-1, 0, 1}
	for _, dx := range delta {
		for _, dy := range delta {
			for _, dz := range delta {
				if !(dx == 0 && dy == 0 && dz == 0) {
					points = append(points, Point{p.X + dx, p.Y + dy, p.Z + dz, p.W})
				}
			}
		}
	}

	return points
}

func (p Point) Neighbours4d() (points []Point) {
	delta := []Int{-1, 0, 1}
	for _, dx := range delta {
		for _, dy := range delta {
			for _, dz := range delta {
				for _, dw := range delta {
					if !(dx == 0 && dy == 0 && dz == 0 && dw == 0) {
						points = append(points, Point{p.X + dx, p.Y + dy, p.Z + dz, p.W + dw})
					}
				}
			}
		}
	}

	return points
}

type Pocket = map[Point]bool

func createPocket(input string) Pocket {
	lines := strings.Fields(input)
	sy := len(lines)
	sx := len(lines[0])
	cubes := make(map[Point]bool, sx*sy)
	for x, line := range lines {
		for y, c := range line {
			if string(c) == "#" {
				cubes[Point{Int(x), Int(y), 0, 0}] = true
			}
		}
	}

	return cubes
}

func solve(pocket Pocket, cycles int, neighbours func(p Point) []Point) int {
	deltas := neighbours(Point{0, 0, 0, 0})

	for i := 0; i < cycles; i++ {

		next := Pocket{}
		inactive := make(map[Point]int)
		for point, _ := range pocket {

			active := 0

			for _, d := range deltas {
				n := point.Add(d)
				if pocket[n] {
					active++
				} else {
					inactive[n] += 1
				}
			}

			if active == 2 || active == 3 {
				next[point] = true
			}

		}

		for p, n := range inactive {
			if n == 3 {
				next[p] = true
			}
		}

		pocket = next
	}

	total := 0
	for _, v := range pocket {
		if v {
			total++
		}
	}

	return total
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day17.txt")

	pocket := createPocket(input)

	fmt.Println("(part1)", solve(pocket, 6, func(p Point) []Point { return p.Neighbours3d() })) // 317
	fmt.Println("(part2)", solve(pocket, 6, func(p Point) []Point { return p.Neighbours4d() })) // 1692

}
