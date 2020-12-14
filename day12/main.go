package main

import (
	"adventofcode2020"
	"fmt"
	"math"
	"regexp"
)

type Instruction struct {
	Action rune
	Value  int
}

type Coord struct {
	x int
	y int
}

func parseInstruction(input []string) []Instruction {
	regex := regexp.MustCompile(`([A-Z])(\d+)`)
	instructions := make([]Instruction, len(input))
	for i, l := range input {
		m := regex.FindStringSubmatch(l)
		instructions[i] = Instruction{Action: rune(m[1][0]), Value: adventofcode2020.MustAtoi(m[2])}
	}

	return instructions
}

func rotate(pos Coord, n int, dir int) Coord {
	var rotate func(c Coord) Coord
	if dir > 0 {
		rotate = func(c Coord) Coord { return Coord{x: c.y, y: -c.x} }
	} else {
		rotate = func(c Coord) Coord { return Coord{x: -c.y, y: c.x} }
	}
	for i := 0; i < n; i++ {
		pos = rotate(pos)
	}

	return pos
}

func SolvePart1(instructions []Instruction) int {
	pos := Coord{0, 0}
	dir := Coord{1, 0}

	for _, i := range instructions {
		switch i.Action {
		case 'N':
			pos.y += i.Value
		case 'S':
			pos.y -= i.Value
		case 'E':
			pos.x += i.Value
		case 'W':
			pos.x -= i.Value
		case 'R':
			dir = rotate(dir, i.Value/90, 1)
		case 'L':
			dir = rotate(dir, i.Value/90, -1)
		case 'F':
			pos.x += dir.x * i.Value
			pos.y += dir.y * i.Value
		}
	}

	return int(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
}

func SolvePart2(instructions []Instruction) int {
	pos := Coord{0, 0}
	wp := Coord{10, 1}

	for _, i := range instructions {
		switch i.Action {
		case 'N':
			wp.y += i.Value
		case 'S':
			wp.y -= i.Value
		case 'E':
			wp.x += i.Value
		case 'W':
			wp.x -= i.Value
		case 'R':
			wp = rotate(wp, i.Value/90, 1)
		case 'L':
			wp = rotate(wp, i.Value/90, -1)
		case 'F':
			pos.x += i.Value * wp.x
			pos.y += i.Value * wp.y
		}
	}

	return int(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	lines, _ := adventofcode2020.ReadInputLines("./input/day12.txt")
	instructions := parseInstruction(lines)
	p1 := SolvePart1(instructions)
	fmt.Println("(part1)", p1)
	p2 := SolvePart2(instructions)
	fmt.Println("(part2)", p2)
}
