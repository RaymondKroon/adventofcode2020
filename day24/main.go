package main

import (
	"adventofcode2020/util"
	"fmt"
	"regexp"
	"strings"
)

////go:generate stringer -type=Move // duplicate if using enumer
//go:generate enumer -type=Move -text

type Move int

const (
	e Move = iota + 1
	se
	sw
	w
	nw
	ne
)

type Coord struct {
	x int
	y int
}

func (c *Coord) MoveXY(dx, dy int) Coord {
	return Coord{
		x: c.x + dx,
		y: c.y + dy,
	}
}

func (c *Coord) Move(move Move) Coord {
	switch move {
	case e:
		return c.MoveXY(2, 0)
	case se:
		return c.MoveXY(1, -1)
	case sw:
		return c.MoveXY(-1, -1)
	case w:
		return c.MoveXY(-2, 0)
	case nw:
		return c.MoveXY(-1, 1)
	case ne:
		return c.MoveXY(1, 1)
	default:
		panic("impossible")
	}
}

func (c *Coord) Neigbours() []Coord {
	neighbours := make([]Coord, 0, 6)
	for _, move := range MoveValues() {
		neighbours = append(neighbours, c.Move(move))
	}

	return neighbours
}

func ParseInstructions(input string) [][]Move {
	lines := strings.Split(input, "\n")
	result := make([][]Move, 0, len(lines))

	regex := regexp.MustCompile(`(e|se|sw|w|nw|ne)`)
	for _, line := range lines {
		tileMoves := make([]Move, 0)
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			move, _ := MoveString(match[1])
			tileMoves = append(tileMoves, move)
		}
		result = append(result, tileMoves)
	}

	return result
}

func FlipTiles(instructions [][]Move) (blacks map[Coord]struct{}) {
	blacks = make(map[Coord]struct{}, len(instructions))
	for _, instruction := range instructions {
		tile := Coord{0, 0}
		for _, move := range instruction {
			tile = tile.Move(move)
		}
		if _, exists := blacks[tile]; exists {
			delete(blacks, tile)
		} else {
			blacks[tile] = struct{}{}
		}
	}

	return blacks
}

func DailyFlip(blacks map[Coord]struct{}) map[Coord]struct{} {
	result := make(map[Coord]struct{})
	whitesWithBlackNeighbours := make(map[Coord]int)
	for tile, _ := range blacks {
		count := 0
		for _, nb := range tile.Neigbours() {
			if _, exist := blacks[nb]; exist {
				count += 1
			} else {
				whitesWithBlackNeighbours[nb] += 1
			}
		}
		if !(count == 0 || count > 2) {
			result[tile] = struct{}{}
		}
	}

	for tile, n := range whitesWithBlackNeighbours {
		if n == 2 {
			result[tile] = struct{}{}
		}
	}

	return result
}

func part2(blacks map[Coord]struct{}) int {
	for day := 0; day < 100; day++ {
		blacks = DailyFlip(blacks)
	}

	return len(blacks)
}

func main() {

	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day24.txt")
	instructions := ParseInstructions(input)

	blacks := FlipTiles(instructions)
	fmt.Println("(p1)", len(blacks))
	fmt.Println("(p2)", part2(blacks))
}
