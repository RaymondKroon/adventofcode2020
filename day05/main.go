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
		if string(c) == down {
			result = (result << 1) | 0
		} else {
			result = (result << 1) | 1
		}
	}

	return result
}

func FindElement(input string, upperBound int, lowCase string) int {
	low := 0
	high := upperBound - 1
	stepsize := upperBound

	for _, a := range input {
		stepsize = stepsize / 2
		if string(a) == lowCase {
			high = high - stepsize
		} else {
			low = low + stepsize
		}
	}

	return low
}

func ParseBoardingPassOld(input string) BoardingPass {
	r := input[0:7]
	s := input[7:]

	row := FindElement(r, 128, "F")
	seat := FindElement(s, 8, "L")
	return BoardingPass{
		Row:  row,
		Seat: seat,
		ID:   row*8 + seat,
	}
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
	defer adventofcode2020.Stopwatch("Run")()
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
