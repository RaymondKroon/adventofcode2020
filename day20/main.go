package main

import (
	"adventofcode2020/util"
	"fmt"
	"regexp"
	"strings"
)

var (
	idRegex = regexp.MustCompile(`\d+`)
)

type Tile struct {
	pixels [10][10]rune
}

func ParseTile(input string) (id int, tile Tile) {
	lines := strings.Split(input, "\n")
	id = util.MustAtoi(idRegex.FindString(lines[0]))
	for y, line := range lines[1:] {
		for x, c := range line {
			tile.pixels[y][x] = c
		}
	}
	return
}

func loadTiles(input string) (tiles map[int]Tile) {
	parts := strings.Split(input, "\n\n")
	tiles = make(map[int]Tile, len(parts))
	for _, part := range parts {
		id, tile := ParseTile(part)
		tiles[id] = tile
	}

	return tiles
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day20.txt")

	tiles := loadTiles(input)
	fmt.Println(tiles)

}
