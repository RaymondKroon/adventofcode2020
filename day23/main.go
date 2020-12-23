package main

import (
	"adventofcode2020/util"
	"bytes"
	"container/ring"
	"fmt"
)

type Int = util.Int

type CrabCups struct {
	size      int
	cups      *ring.Ring
	positions map[int]*ring.Ring
}

func (cc *CrabCups) Values() []int {
	values := make([]int, cc.size, cc.size)
	for i := 0; i < cc.size; i++ {
		values[i] = cc.cups.Value.(int)
		cc.cups = cc.cups.Next()
	}

	return values
}

func (cc *CrabCups) Move() (self *CrabCups) {

	current := cc.cups.Value.(int)
	three := cc.cups.Unlink(3)

	picked := map[int]*struct{}{}

	for i := 0; i < 3; i++ {
		picked[three.Value.(int)] = &struct{}{}
		three = three.Next()
	}

	target := current - 1
	for picked[target] != nil {
		target -= 1
	}
	if target == 0 {
		target = cc.size
		for picked[target] != nil {
			target -= 1
		}
	}

	cc.positions[target].Link(three)
	cc.cups = cc.cups.Next()

	return cc
}

func NewCrabCups(initCups []int, totalCups int) CrabCups {
	cups := ring.New(totalCups)
	positions := map[int]*ring.Ring{}
	for i := 0; i < totalCups; i++ {
		if i < len(initCups) {
			cups.Value = initCups[i]
		} else {
			cups.Value = i + 1
		}
		positions[cups.Value.(int)] = cups
		cups = cups.Next()
	}

	return CrabCups{
		size:      totalCups,
		cups:      cups,
		positions: positions,
	}
}

func NewCrabCupsFromString(input string, nCups int) CrabCups {
	cupValues := make([]int, len(input))
	for i, c := range input {
		cupValues[i] = util.MustAtoi(string(c))
	}

	return NewCrabCups(cupValues, nCups)
}

func part1(cc CrabCups, moves int) string {
	for i := 0; i < moves; i++ {
		cc.Move()
	}

	var buf bytes.Buffer
	start := cc.positions[1]
	for i := 0; i < cc.size-1; i++ {
		start = start.Next()
		buf.WriteString(fmt.Sprint(start.Value.(int)))
	}

	return buf.String()
}

func part2(cc CrabCups, moves int) int {
	for i := 0; i < moves; i++ {
		cc.Move()
	}

	one := cc.positions[1]
	return one.Move(1).Value.(int) * one.Move(2).Value.(int)
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day23.txt")

	fmt.Println("(p1)", part1(NewCrabCupsFromString(input, 9), 100))            //27956483
	fmt.Println("(p2)", part2(NewCrabCupsFromString(input, 1000000), 10000000)) //18930983775
}
