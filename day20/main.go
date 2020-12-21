package main

import (
	"adventofcode2020/util"
	"fmt"
	"math"
	"regexp"
	"strings"
)

////go:generate genny -in=../util/queue.go -out=gen-$GOFILE -pkg=main gen "ValueType=Coord"
//go:generate genny -in=../util/slice.go -out=gen-slice-$GOFILE -pkg=main gen "ValueType=Coord,Tile"

var (
	idRegex    = regexp.MustCompile(`\d+`)
	Directions = [4]Direction{Top, Right, Bottom, Left}
)

type Direction uint8

const (
	Top Direction = iota + 1
	Right
	Bottom
	Left
)

const XSize = 10
const YSize = 10

type PixelsType = [XSize][YSize]rune

type Tile struct {
	Id     int
	pixels PixelsType
}

func (t *Tile) RotateCW() *Tile {
	var rotated = PixelsType{}
	for y := 0; y < YSize; y++ {
		for x := 0; x < XSize; x++ {
			rotated[x][YSize-y-1] = t.pixels[y][x]
		}
	}

	return &Tile{pixels: rotated, Id: t.Id}
}

func (t *Tile) FlipH() *Tile {
	var flipped = PixelsType{}
	for y := 0; y < YSize; y++ {
		for x := 0; x < XSize; x++ {
			flipped[y][XSize-x-1] = t.pixels[y][x]
		}
	}

	return &Tile{pixels: flipped, Id: t.Id}
}

func (t *Tile) GetAllPositions() []Tile {
	positions := make([]Tile, 8)

	for f := 0; f <= 1; f++ {
		var p *Tile
		if f != 0 {
			p = t.FlipH()
		} else {
			p = t
		}
		for r := 0; r < 4; r++ {
			positions[r+f*4] = *p
			p = p.RotateCW()
		}
	}

	return positions
}

func (t *Tile) IsTopNeighbour(o *Tile) bool {
	return t.pixels[0] == o.pixels[YSize-1]
}

func (t *Tile) IsRightNeighbour(o *Tile) bool {
	for y := 0; y < YSize; y++ {
		if t.pixels[y][XSize-1] != o.pixels[y][0] {
			return false
		}
	}

	return true
}

func (t *Tile) IsBottomNeighbour(o *Tile) bool {
	return o.IsTopNeighbour(t)
}

func (t *Tile) IsLeftNeighbour(o *Tile) bool {
	return o.IsRightNeighbour(t)
}

func ParseTile(lines []string) (tile Tile) {
	for y, line := range lines {
		for x, c := range line {
			tile.pixels[y][x] = c
		}
	}

	return
}

func ParseTileWithTitle(input string) (id int, tile Tile) {
	lines := strings.Split(input, "\n")
	id = util.MustAtoi(idRegex.FindString(lines[0]))
	tile = ParseTile(lines[1:])
	tile.Id = id
	return
}

func loadTiles(input string) (tiles []Tile) {
	parts := strings.Split(input, "\n\n")
	tiles = make([]Tile, len(parts))
	for i, part := range parts {
		_, tile := ParseTileWithTitle(part)
		tiles[i] = tile
	}

	return tiles
}

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

func CloneImage(image map[Coord]Tile) map[Coord]Tile {
	clone := make(map[Coord]Tile, len(image))
	for k, v := range image {
		clone[k] = v
	}

	return clone
}

func solve(tiles []Tile, coords []Coord, partialImage map[Coord]Tile, targetSize int) (result bool, image map[Coord]Tile) {
	if len(tiles) == 0 && len(partialImage) == targetSize {
		return true, partialImage
	}

	if len(coords) == 0 {
		return false, nil
	}

	coord := coords[0]
	var tile Tile
	for i := 0; i < len(tiles); i++ {
		tile = tiles[i]
		//for _, permutatedTile := range []Tile{tile} {
		for _, permutatedTile := range tile.GetAllPositions() {
			// top
			if nbTile, exists := partialImage[coord.AddValue(Coord{1, 0})]; exists {
				if !permutatedTile.IsTopNeighbour(&nbTile) {
					continue
				}
			}
			// right
			if nbTile, exists := partialImage[coord.AddValue(Coord{0, 1})]; exists {
				if !permutatedTile.IsRightNeighbour(&nbTile) {
					continue
				}
			}

			// bottom
			if nbTile, exists := partialImage[coord.AddValue(Coord{-1, 0})]; exists {
				if !permutatedTile.IsBottomNeighbour(&nbTile) {
					continue
				}
			}

			// left
			if nbTile, exists := partialImage[coord.AddValue(Coord{0, -1})]; exists {
				if !permutatedTile.IsLeftNeighbour(&nbTile) {
					continue
				}
			}

			newPartialImage := CloneImage(partialImage)
			newPartialImage[coord] = permutatedTile

			if ok, image := solve(RemoveFromTileSlice(tiles, i), RemoveFromCoordSlice(coords, 0), newPartialImage, targetSize); ok {
				return ok, image
			}
		}
	}

	return false, nil
}

func part1(tiles *[]Tile) int {
	side := int(math.Sqrt(float64(len(*tiles))))
	if side*side != len(*tiles) {
		panic("NOT SQUARE")
	}

	coords := make([]Coord, len(*tiles))
	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			coords[x+side*y] = Coord{x, y}
		}
	}

	ok, image := solve(*tiles, coords, map[Coord]Tile{}, len(*tiles))
	if ok {
		return image[Coord{0, 0}].Id *
			image[Coord{0, side - 1}].Id *
			image[Coord{side - 1, 0}].Id *
			image[Coord{side - 1, side - 1}].Id
	} else {
		return -1
	}
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day20.txt")

	tiles := loadTiles(input)

	fmt.Println("(p1)", part1(&tiles))

}
