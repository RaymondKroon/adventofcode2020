package main

import (
	"adventofcode2020"
	"fmt"
	"regexp"
)

type Instruction struct {
	Op  string
	Arg int
}

func ParseInstructions(input []string) []Instruction {
	regex := regexp.MustCompile(`([a-z]{3}) ([+|-]\d+)`)
	result := make([]Instruction, len(input))
	for i, line := range input {
		match := regex.FindStringSubmatch(line)
		result[i] = Instruction{
			Op:  match[1],
			Arg: adventofcode2020.MustAtoi(match[2]),
		}
	}

	return result
}

type Exit struct {
	Acc int
	Ok  bool
}

func RunUntilLoopOrNormalExit(instructions []Instruction) Exit {
	acc := 0
	index := 0
	visited := make(map[int]bool)

	for found := false; !found && index < len(instructions); found = visited[index] {
		visited[index] = true
		instruction := instructions[index]
		switch instruction.Op {
		case "nop":
			index++
			break
		case "acc":
			index++
			acc += instruction.Arg
			break
		case "jmp":
			index += instruction.Arg
			break
		default:
			break
		}
	}

	return Exit{
		Acc: acc,
		Ok:  index >= len(instructions),
	}
}

func Repair(instructions []Instruction) int {
	origOp := ""
	for index := 0; index < len(instructions); index++ {
		instruction := &instructions[index]
		if instruction.Op == "nop" {
			origOp = "nop"
			instruction.Op = "jmp"
		} else if instruction.Op == "jmp" {
			origOp = "jmp"
			instruction.Op = "nop"
		}

		if instruction.Op != "acc" {
			exit := RunUntilLoopOrNormalExit(instructions)
			if exit.Ok {
				return exit.Acc
			} else {
				instruction.Op = origOp
			}
		}
	}

	return -1
}

func main() {
	instructionLines, _ := adventofcode2020.ReadInputLines("./input/day08.txt")
	instructions := ParseInstructions(instructionLines)
	//for _, i := range instructions {
	//    fmt.Println(i)
	//}

	fmt.Printf("(part1) Acc: %d\n", RunUntilLoopOrNormalExit(instructions).Acc)
	fmt.Printf("(part2) Acc: %d\n", Repair(instructions))
}
