package main

import (
	"adventofcode2020"
	"fmt"
	"sort"
)

type BoardingPass struct {
	Row  int
	Seat int
	ID   int
}

func ParseBinaryCode(code string, down string, up string) int {
	result := 0
	for _, c := range code {
		var d int
		if string(c) == down {
			d = 0
		} else {
			d = 1
		}
		result = (result << 1) | d
	}

	return result
}

func ParseBoardingPass(input string) BoardingPass {
	r := input[0:7]
	s := input[7:]

	row := ParseBinaryCode(r, "F", "B")
	seat := ParseBinaryCode(s, "L", "R")
	return BoardingPass{
		Row:  row,
		Seat: seat,
		ID:   row*8 + seat,
	}
}

func main() {
	stringInput, _ := adventofcode2020.ReadInputLines("./input/day05.txt")
	var bps []BoardingPass
	for _, input := range stringInput {
		bps = append(bps, ParseBoardingPass(input))
	}

	maxId := 0
	var IDs []int
	for _, bp := range bps {
		if bp.ID > maxId {
			maxId = bp.ID
		}
		IDs = append(IDs, bp.ID)
	}

	fmt.Printf("(part1) maxId: %d\n", maxId)

	sort.Ints(IDs)
	for j, _ := range IDs[:len(IDs)-1] {

		if IDs[j]+1 != IDs[j+1] {
			fmt.Printf("(part1) mySeat: %d\n", IDs[j]+1)
		}
	}
}
