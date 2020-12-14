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

type FloorPlan struct {
	nRows int
	nCols int
	inner [][]Cell
}

func (fp *FloorPlan) Get(row int, col int) (Cell, bool) {
	if row < 0 || col < 0 || row >= fp.nRows || col >= fp.nCols {
		return Empty, false
	} else {
		return fp.inner[row][col], true
	}
}

func solve(plan *FloorPlan, occupiedLimit int, neighbourFinder func(plan *FloorPlan, coord Coord, delta Coord) Coord) int {
	delta := []Coord{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for changed := true; changed; {
		changed = false
		clonedInner := make([][]Cell, len(plan.inner))
		for i := range plan.inner {
			clonedInner[i] = make([]Cell, len(plan.inner[i]))
		}
		next := FloorPlan{nRows: plan.nRows, nCols: plan.nCols, inner: clonedInner}

		for r := 0; r < plan.nRows; r++ {
			for c := 0; c < plan.nCols; c++ {

				coord := Coord{row: r, col: c}
				cell, _ := plan.Get(r, c)

				next.inner[r][c] = cell
				if cell == Floor {
					continue
				}

				occupiedNeighbours := 0
				for _, d := range delta {
					neighbourCoord := neighbourFinder(plan, coord, d)
					if n, ok := plan.Get(neighbourCoord.row, neighbourCoord.col); ok && n == Occupied {
						occupiedNeighbours++
						if occupiedNeighbours >= occupiedLimit {
							break
						}
					}
				}

				if cell == Occupied && occupiedNeighbours >= occupiedLimit {
					next.inner[r][c] = Empty
					changed = true
					continue
				}
				if cell == Empty && occupiedNeighbours == 0 {
					next.inner[r][c] = Occupied
					changed = true
					continue
				}
			}
		}

		plan = &next
	}

	occupied := 0
	for _, row := range plan.inner {
		for _, cell := range row {
			if cell == Occupied {
				occupied += 1
			}
		}
	}

	return occupied
}

func solvePart1(plan *FloorPlan) (occupied int) {
	return solve(plan, 4, func(plan *FloorPlan, coord Coord, delta Coord) Coord { return coord.AddValue(delta) })
}

func solvePart2(plan *FloorPlan) (occupied int) {
	return solve(plan, 5, func(plan *FloorPlan, coord Coord, delta Coord) Coord {
		newCoord := coord.AddValue(delta)
		for n, ok := plan.Get(newCoord.row, newCoord.col); ok && n == Floor; n, ok = plan.Get(newCoord.row, newCoord.col) {
			newCoord = newCoord.AddValue(delta)
		}
		return newCoord
	})
}

func parseFloorplan(input []string) FloorPlan {
	floorPlan := make([][]Cell, len(input))
	for r, line := range input {
		row := make([]Cell, len(line))
		for c, col := range line {
			var t Cell
			switch col {
			case '.':
				t = Floor
			case 'L':
				t = Empty
			}
			row[c] = t
		}
		floorPlan[r] = row
	}
	return FloorPlan{
		nRows: len(floorPlan),
		nCols: len(floorPlan[0]),
		inner: floorPlan,
	}
}

//func printFloorPlan(fp FloorPlan, nRow int, nCol int) {
//	for row := 0; row < nRow; row++ {
//		for col := 0; col < nCol; col++ {
//			fmt.Printf("%s", fp[Coord{row, col}].String())
//		}
//		fmt.Print("\n")
//	}
//
//	fmt.Println("----------")
//}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	lines, _ := adventofcode2020.ReadInputLines("./input/day11.txt")
	floorplan := parseFloorplan(lines)
	p1 := solvePart1(&floorplan) //2424
	fmt.Println("(part1)", p1)

	p2 := solvePart2(&floorplan) //2208
	fmt.Println("(part2)", p2)
}
