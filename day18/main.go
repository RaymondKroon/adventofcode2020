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

//https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func parse(expressionStr string, precedes func(op1, op2 string) bool) (expression []string) {
	tokenMatcher := regexp.MustCompile(`([0-9\(\)\+\*])`)
	matches := tokenMatcher.FindAllStringSubmatch(expressionStr, -1)
	outputStack := util.NewStringStack()
	operatorStack := util.NewStringStack()

	for _, match := range matches {
		token := match[0]
		switch token {
		case "+":
			for op := operatorStack.Peek(); op != nil && *op != "(" && precedes(*op, token); op = operatorStack.Peek() {
				outputStack.Push(*operatorStack.Pop())
			}
			operatorStack.Push(token)
		case "*":
			for op := operatorStack.Peek(); op != nil && *op != "(" && precedes(*op, token); op = operatorStack.Peek() {
				outputStack.Push(*operatorStack.Pop())
			}
			operatorStack.Push(token)
		case "(":
			operatorStack.Push(token)
		case ")":
			for t := operatorStack.Pop(); t != nil && *t != "("; t = operatorStack.Pop() {
				outputStack.Push(*t)
			}
		default:
			outputStack.Push(token)
		}
	}

	for op := operatorStack.Pop(); op != nil; op = operatorStack.Pop() {
		outputStack.Push(*op)
	}
	return outputStack.ToArray()
}

func evaluate(expression []string) int {
	stack := util.NewIntStack()

	for _, e := range expression {
		switch e {
		case "+":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(*a + *b)
		case "*":
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(*a * *b)
		default:
			stack.Push(util.MustAtoi(e))
		}
	}

	result := stack.Pop()
	return *result
}

func linear(op1, op2 string) bool {
	return true
}

func addFirst(op1, op2 string) bool {
	if op1 == "+" {
		return true
	} else {
		return false
	}
}

func solvePart1(expressions []string) int {
	total := 0
	for _, exp := range expressions {
		total += evaluate(parse(exp, linear))
	}

	return total
}

func solvePart2(expressions []string) int {
	total := 0
	for _, exp := range expressions {
		total += evaluate(parse(exp, addFirst))
	}

	return total
}

func main() {
	defer util.Stopwatch("Run")()
	expressions, _ := util.ReadInputLines("./input/day18.txt")
	fmt.Println("(part1)", solvePart1(expressions)) //7147789965219
	fmt.Println("(part2)", solvePart2(expressions)) //136824720421264

}
