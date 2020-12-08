package main

import (
	"adventofcode2020"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var policyRegex *regexp.Regexp

func init() {
	policyRegex = regexp.MustCompile(`(\d+)-(\d+) ([a-z])`)
}

type Policy struct {
	char string
	min  int
	max  int
}

func ParsePolicy(s string) (Policy, error) {
	match := policyRegex.FindAllStringSubmatch(s, 1)[0]
	min, _ := strconv.Atoi(match[1])
	max, _ := strconv.Atoi(match[2])
	return Policy{
		char: match[3],
		min:  min,
		max:  max,
	}, nil
}

type PasswordLine struct {
	policy Policy
	input  string
}

func ParseInput(in []string) ([]PasswordLine, error) {
	var result []PasswordLine
	for _, s := range in {
		parts := strings.SplitN(s, ":", 2)
		policy, _ := ParsePolicy(parts[0])
		line := PasswordLine{
			policy: policy,
			input:  strings.TrimSpace(parts[1]),
		}
		result = append(result, line)
	}

	return result, nil
}

func Part1CountCorrectPasswords(input []PasswordLine) int {
	result := 0
	for _, l := range input {
		counted := strings.Count(l.input, l.policy.char)
		if counted >= l.policy.min && counted <= l.policy.max {
			result++
		}
	}
	return result
}

func Part2CountCorrectPasswords(input []PasswordLine) int {
	result := 0
	for _, l := range input {
		first := false
		if len(l.input) > l.policy.min-1 {
			first = string(l.input[l.policy.min-1]) == l.policy.char
		}
		second := false
		if len(l.input) > l.policy.max-1 {
			second = string(l.input[l.policy.max-1]) == l.policy.char
		}
		if first != second {
			result++
		}
	}
	return result
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	stringInput, _ := adventofcode2020.ReadInputLines("./input/day02.txt")
	input, _ := ParseInput(stringInput)
	fmt.Printf("(part1) Correct password: %d\n", Part1CountCorrectPasswords(input))
	fmt.Printf("(part2) Correct password: %d\n", Part2CountCorrectPasswords(input))
}
