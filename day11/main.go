package main

import (
	"adventofcode2020"
	"fmt"
)

type Coord struct {
	row int
	col int
}

func (c *Coord) AddValue(other Coord) Coord {
	return Coord{
		row: c.row + other.row,
		col: c.col + other.col,
	}
}

type Cell uint8

const (
	Floor Cell = iota + 1
	Empty
	Occupied
)

func (c Cell) String() string {
	switch c {
	case Floor:
		return "."
	case Empty:
		return "L"
	case Occupied:
		return "#"
	default:
		return "X"
	}
}

type FloorPlan = map[Coord]Cell

func solve(plan FloorPlan, occupiedLimit int, neighbourFinder func(plan FloorPlan, coord Coord, delta Coord) Coord) int {
	delta := []Coord{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for changed := true; changed; {
		changed = false
		next := FloorPlan{}

		for coord, cell := range plan {
			occupiedNeighbours := 0
			for _, d := range delta {
				if plan[neighbourFinder(plan, coord, d)] == Occupied {
					occupiedNeighbours++
				}
			}

			if cell == Occupied && occupiedNeighbours >= occupiedLimit {
				next[coord] = Empty
				changed = true
				continue
			}
			if cell == Empty && occupiedNeighbours == 0 {
				next[coord] = Occupied
				changed = true
				continue
			}

			next[coord] = cell
		}

		plan = next
	}

	occupied := 0
	for _, v := range plan {
		if v == Occupied {
			occupied += 1
		}
	}

	return occupied
}

func solvePart1(plan FloorPlan) (occupied int) {
	return solve(plan, 4, func(plan FloorPlan, coord Coord, delta Coord) Coord { return coord.AddValue(delta) })
}

func solvePart2(plan FloorPlan) (occupied int) {
	return solve(plan, 5, func(plan FloorPlan, coord Coord, delta Coord) Coord {
		newCoord := coord.AddValue(delta)
		for plan[newCoord] == Floor {
			newCoord = newCoord.AddValue(delta)
		}
		return newCoord
	})
}

func parseFloorplan(input []string) FloorPlan {
	floorPlan := FloorPlan{}
	for r, line := range input {
		for c, col := range line {
			var t Cell
			switch col {
			case '.':
				t = Floor
				break
			case 'L':
				t = Empty
			}
			floorPlan[Coord{row: r, col: c}] = t
		}
	}
	return floorPlan
}

func printFloorPlan(fp FloorPlan, nRow int, nCol int) {
	for row := 0; row < nRow; row++ {
		for col := 0; col < nCol; col++ {
			fmt.Printf("%s", fp[Coord{row, col}].String())
		}
		fmt.Print("\n")
	}

	fmt.Println("----------")
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	lines, _ := adventofcode2020.ReadInputLines("./input/day11.txt")
	floorplan := parseFloorplan(lines)
	p1 := solvePart1(floorplan) //2424
	fmt.Println("(part1)", p1)

	p2 := solvePart2(floorplan)
	fmt.Println("(part2)", p2)
}
