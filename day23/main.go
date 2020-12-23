package main

import (
	"adventofcode2020/util"
	"fmt"
)

const Size = 9

type Int = util.Int

type CrabCups [Size]Int

func (cc CrabCups) Move() CrabCups {
	current := cc[0]
	three := cc[1:4]
	remainder := cc[4:]
	target := current - 1
	var destination int
	ok := false
	for !ok && target > 0 {
		ok, destination = util.IntInSlice(target, remainder)
		target -= 1
	}
	if !ok {
		_, destination = util.MaxInt(remainder)
	}

	result := make([]Int, 0, 9)
	result = append(result, remainder[:destination+1]...)
	result = append(result, three...)
	if destination < len(remainder) {
		result = append(result, remainder[destination+1:]...)
	}
	result = append(result, current)

	var newCC CrabCups
	copy(newCC[:], result)
	return newCC
}

func NewCrabCups(input string) CrabCups {
	var cups CrabCups
	for i := 0; i < Size; i++ {
		cups[i] = Int(util.MustAtoi(string(input[i])))
	}

	return cups
}

func part1(cc CrabCups, moves int) string {
	for i := 0; i < moves; i++ {
		cc = cc.Move()
	}

	var one int
	for i, val := range cc {
		if val == 1 {
			one = i
			break
		}
	}
	var result []Int
	result = append(result, cc[one+1:]...)
	result = append(result, cc[:one]...)
	return util.IntJoint(result, "")
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day23.txt")

	game := NewCrabCups(input)

	fmt.Println("(p1)", part1(game, 100))
}
