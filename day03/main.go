package main

import (
	"adventofcode2020"
	"fmt"
)

type Map interface {
	Move(right int, down int)
	MapSquare() string
	IsDown() bool
	Reset()
}

type MapImpl struct {
	inner [][]string
	x     int
	y     int
	xSize int
	ySize int
}

func (m *MapImpl) Move(right int, down int) {
	if !m.IsDown() {
		m.x = (m.x + right) % m.xSize
		m.y = m.y + down
	}
}

func (m *MapImpl) MapSquare() string {
	return m.inner[m.y][m.x]
}

func (m *MapImpl) IsDown() bool {
	return m.y == m.ySize-1
}

func (m *MapImpl) Reset() {
	m.x = 0
	m.y = 0
}

func CreateMap(input []string) Map {
	var result [][]string
	for _, in := range input {
		var line []string
		for _, c := range in {
			line = append(line, string(c))
		}
		result = append(result, line)
	}

	return &MapImpl{
		inner: result,
		x:     0,
		y:     0,
		xSize: len(result[0]),
		ySize: len(result),
	}
}

func MoveDownAndCountTrees(stepsizeRight int, stepsizeDown int, m Map) int {
	nTrees := 0
	for !m.IsDown() {
		m.Move(stepsizeRight, stepsizeDown)
		if m.MapSquare() == "#" {
			nTrees++
		}
	}

	return nTrees
}

func Part1(m Map) int {
	return MoveDownAndCountTrees(3, 1, m)
}

func Part2(m Map) int {
	steps := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	multiplied := 1
	for _, step := range steps {
		multiplied *= MoveDownAndCountTrees(step[0], step[1], m)
		m.Reset()
	}

	return multiplied
}

func main() {
	stringInput, _ := adventofcode2020.ReadInput("./input/day03.txt")
	m := CreateMap(stringInput)

	fmt.Printf("(part1) nTrees: %d\n", Part1(m))
	m.Reset()
	fmt.Printf("(part2) Result: %d\n", Part2(m))
}
