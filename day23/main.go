package main

import (
	"adventofcode2020/util"
	"bytes"
	"fmt"
)

type Int = util.Int

type CrabCups struct {
	size int
	cur  int
	cups []int
}

func (cc *CrabCups) Values() []int {
	values := make([]int, cc.size, cc.size)
	p := cc.cur
	for i := 0; i < cc.size; i++ {
		values[i] = p
		p = cc.cups[p]
	}

	return values
}

func (cc *CrabCups) Move() (self *CrabCups) {
	remove1 := cc.cups[cc.cur]
	remove2 := cc.cups[remove1]
	remove3 := cc.cups[remove2]

	target := cc.cur
	for target == cc.cur || target == remove1 || target == remove2 || target == remove3 {
		target -= 1
		if target == 0 {
			target = cc.size
		}
	}

	insert_point := cc.cups[target]
	cc.cups[cc.cur] = cc.cups[remove3]
	cc.cups[target] = remove1

	cc.cups[remove3] = insert_point

	cc.cur = cc.cups[cc.cur]

	return cc
}

func NewCrabCups(initCups []int, totalCups int) CrabCups {
	cups := make([]int, totalCups)
	for i := 0; i < totalCups; i++ {
		if i < len(initCups) {
			cups[i] = initCups[i]
		} else {
			cups[i] = i + 1
		}
	}
	arr := make([]int, totalCups+1)
	next := append(cups[1:], cups[0])
	for i := 0; i < totalCups; i++ {
		arr[cups[i]] = next[i]
	}

	return CrabCups{
		size: totalCups,
		cups: arr,
		cur:  initCups[0],
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
	v := cc.cups[1]
	for i := 0; i < cc.size-1; i++ {
		buf.WriteString(fmt.Sprint(v))
		v = cc.cups[v]
	}

	return buf.String()
}

func part2(cc CrabCups, moves int) int {
	for i := 0; i < moves; i++ {
		cc.Move()
	}

	a := cc.cups[1]
	b := cc.cups[a]
	return a * b
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day23.txt")

	fmt.Println("(p1)", part1(NewCrabCupsFromString(input, 9), 100))            //27956483
	fmt.Println("(p2)", part2(NewCrabCupsFromString(input, 1000000), 10000000)) //18930983775
}
