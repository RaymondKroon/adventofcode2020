package main

import (
	"adventofcode2020/util"
	"fmt"
	"strings"
)

type Point struct {
	X int
	Y int
	Z int
	W int
}

func (p Point) Neighbours3d() (points []Point) {
	delta := []int{-1, 0, 1}
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
	delta := []int{-1, 0, 1}
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
			var active bool
			if string(c) == "#" {
				active = true
			} else {
				active = false
			}
			cubes[Point{x, y, 0, 0}] = active
		}
	}

	return cubes
}

func solve(pocket Pocket, cycles int, neighbours func(p Point) []Point) int {
	for i := 0; i < cycles; i++ {
		// increase Pocket
		for point, _ := range pocket {
			for _, n := range neighbours(point) {
				pocket[n] = pocket[n]
			}
		}

		new := Pocket{}
		for point, _ := range pocket {
			activeNeighbours := 0
			for _, n := range neighbours(point) {
				if pocket[n] {
					activeNeighbours++
				}
			}

			if pocket[point] {
				if activeNeighbours == 2 || activeNeighbours == 3 {
					new[point] = true
				}
			} else {
				if activeNeighbours == 3 {
					new[point] = true
				}
			}

		}

		pocket = new
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

	fmt.Println("(part1)", solve(pocket, 6, func(p Point) []Point { return p.Neighbours3d() }))
	fmt.Println("(part2)", solve(pocket, 6, func(p Point) []Point { return p.Neighbours4d() }))
}
