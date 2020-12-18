package main

import (
	"adventofcode2020/util"
	"fmt"
	"regexp"
)

type Operation = func(a, b int) int

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

type StackValue struct {
	Value int
	Op    Operation
}

type Stack = []*StackValue

func popStack(stack Stack) (val *StackValue, poppedStack Stack) {
	prev := len(stack) - 1
	val = stack[prev]
	poppedStack = stack[:prev]

	return
}

func evaluate(expression string) int {
	tokenMatcher := regexp.MustCompile(`([0-9\(\)\+\*])`)
	match := tokenMatcher.FindAllStringSubmatch(expression, -1)

	stack := Stack{}

	cur := &StackValue{Value: 0, Op: Add}

	for _, token := range match {
		switch token[0] {
		case "+":
			cur.Op = Add
		case "*":
			cur.Op = Multiply
		case "(":
			stack = append(stack, cur)
			cur = &StackValue{0, Add}
		case ")":
			subVal := cur.Value
			cur, stack = popStack(stack)
			cur.Value = cur.Op(cur.Value, subVal)
		default:
			val := util.MustAtoi(token[0])
			cur.Value = cur.Op(cur.Value, val)
		}
	}

	return cur.Value
}

func solvePart1(expressions []string) int {
	total := 0
	for _, exp := range expressions {
		total += evaluate(exp)
	}

	return total
}

func main() {
	defer util.Stopwatch("Run")()
	expressions, _ := util.ReadInputLines("./input/day18.txt")
	fmt.Println("(part1)", solvePart1(expressions))
}
