package main

import (
	"adventofcode2020/util"
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strings"
)

////go:generate genny -in=../util/queue.go -out=gen-$GOFILE -pkg=main gen "ValueType=Coord"
//go:generate genny -in=../util/slice.go -out=gen-slice-$GOFILE -pkg=main gen "SliceType=Coord,Tile"

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

type PixelsType = [][]rune

type Tile struct {
	Id       int
	pixels   PixelsType
	sideSize int
}

func (t *Tile) String() string {
	var buf bytes.Buffer
	first := true
	buf.WriteString(fmt.Sprintf("Tile %d:\n", t.Id))
	for _, line := range t.pixels {
		if !first {
			buf.WriteString("\n")
		}
		for _, c := range line {
			buf.WriteString(string(c))
		}
		first = false
	}

	return buf.String()
}

func (t Tile) Equals(o Tile) bool {
	if t.sideSize != o.sideSize {
		return false
	}
	for y := 0; y < t.sideSize; y++ {
		for x := 0; x < t.sideSize; x++ {
			if t.pixels[y][x] != o.pixels[y][x] {
				return false
			}
		}
	}

	return true
}

func (t *Tile) RotateCW() *Tile {
	rotated := make([][]rune, t.sideSize)
	for y := 0; y < t.sideSize; y++ {
		rotated[y] = make([]rune, t.sideSize)
	}

	for y := range t.pixels {
		for x := range t.pixels[y] {
			rotated[x][t.sideSize-y-1] = t.pixels[y][x]
		}
	}

	return &Tile{pixels: rotated, Id: t.Id, sideSize: t.sideSize}
}

func (t *Tile) FlipH() *Tile {
	flipped := make([][]rune, t.sideSize)
	for y := 0; y < t.sideSize; y++ {
		flipped[y] = make([]rune, t.sideSize)
	}
	for y := range t.pixels {
		for x := range t.pixels[y] {
			flipped[y][t.sideSize-x-1] = t.pixels[y][x]
		}
	}

	return &Tile{pixels: flipped, Id: t.Id, sideSize: t.sideSize}
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
			if r != 3 {
				p = p.RotateCW()
			}
		}
	}

	return positions
}

func (t *Tile) IsTopNeighbour(o *Tile) bool {
	if t.sideSize != o.sideSize {
		return false
	}

	for x := 0; x < t.sideSize; x++ {
		if t.pixels[0][x] != o.pixels[t.sideSize-1][x] {
			return false
		}
	}

	return true
}

func (t *Tile) IsRightNeighbour(o *Tile) bool {
	if t.sideSize != o.sideSize {
		return false
	}

	for y := 0; y < t.sideSize; y++ {
		if t.pixels[y][t.sideSize-1] != o.pixels[y][0] {
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
	pixels := make([][]rune, len(lines))
	for y, line := range lines {
		pixels[y] = make([]rune, len(line))
		for x, c := range line {
			pixels[y][x] = c
		}
	}

	tile.pixels = pixels
	tile.sideSize = len(lines)

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

func (c Coord) Equals(o Coord) bool {
	return c.col == o.col && c.row == o.row
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
			if nbTile, exists := partialImage[coord.AddValue(Coord{-1, 0})]; exists {
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
			if nbTile, exists := partialImage[coord.AddValue(Coord{+1, 0})]; exists {
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

func SolveImage(tiles *[]Tile) (map[Coord]Tile, int) {
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

	_, image := solve(*tiles, coords, map[Coord]Tile{}, len(*tiles))
	return image, side
}

func part1(image map[Coord]Tile, side int) int {

	//for row := 0; row < side; row++ {
	//	for col := 0; col < side; col++ {
	//		tile := image[Coord{row, col}]
	//		fmt.Println(tile.String())
	//		fmt.Println("-----")
	//	}
	//}

	return image[Coord{0, 0}].Id *
		image[Coord{0, side - 1}].Id *
		image[Coord{side - 1, 0}].Id *
		image[Coord{side - 1, side - 1}].Id
}

func part2(image map[Coord]Tile, side int) int {
	sideSize := side * 8
	completePixels := make([][]rune, sideSize)
	for y := 0; y < sideSize; y++ {
		completePixels[y] = make([]rune, sideSize)
	}

	for row := 0; row < side; row++ {
		for col := 0; col < side; col++ {
			tile := image[Coord{row, col}]
			for y := 0; y < 8; y++ {
				for x := 0; x < 8; x++ {
					completePixels[y+row*8][x+col*8] = tile.pixels[y+1][x+1]
				}
			}
		}
	}

	complete := Tile{
		Id:       0,
		pixels:   completePixels,
		sideSize: sideSize,
	}

	monster :=
		`                  # 
#    ##    ##    ###
 #  #  #  #  #  #`

	monsterCoords := map[Coord]struct{}{}
	for row, line := range strings.Split(monster, "\n") {
		for col, c := range line {
			if c == '#' {
				monsterCoords[Coord{row, col}] = struct{}{}
			}
		}
	}

	for _, p := range complete.GetAllPositions() {

		found := false

		for row := range p.pixels {
			for col := range p.pixels[row] {
				coord := Coord{row, col}
				match := true
				for mCoord := range monsterCoords {
					checkCoord := coord.AddValue(mCoord)
					if checkCoord.row >= complete.sideSize || checkCoord.col >= complete.sideSize || p.pixels[checkCoord.row][checkCoord.col] != '#' {
						match = false
						break
					}
				}

				if match {
					found = true
					for mCoord := range monsterCoords {
						checkCoord := coord.AddValue(mCoord)
						p.pixels[checkCoord.row][checkCoord.col] = 'O'
					}

				}
			}
		}

		if found {
			total := 0
			for y := range p.pixels {
				for x := range p.pixels[y] {
					if p.pixels[y][x] == '#' {
						total += 1
					}
				}
			}

			fmt.Println(p.String())

			return total
		}

	}

	return -1
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day20.txt")

	tiles := loadTiles(input)
	image, side := SolveImage(&tiles)

	fmt.Println("(p1)", part1(image, side)) //7492183537913
	fmt.Println("(p2)", part2(image, side)) //2323

}
