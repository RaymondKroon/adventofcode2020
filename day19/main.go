package main

import (
	"adventofcode2020/util"
	"fmt"
	"regexp"
	"strings"
)

type Rule interface {
	String() string
}

type Literal struct {
	char string
}

type Sequence struct {
	sequence []int
}

type Choice struct {
	left  Sequence
	right Sequence
}

func (l *Literal) String() string {
	return fmt.Sprintf("\"%s\"", l.char)
}

func (s *Sequence) String() string {
	return strings.Join(util.MapIntsToStrings(s.sequence), " ")
}

func NewSequenceFromString(s string) *Sequence {
	splitted := strings.Fields(s)
	seq := make([]int, len(splitted))

	for i, m := range splitted {
		seq[i] = util.MustAtoi(m)
	}

	return &Sequence{sequence: seq}
}

func (c *Choice) String() string {
	return fmt.Sprintf("%s | %s", c.left.String(), c.right.String())
}

var (
	literalRe = regexp.MustCompile(`"([a-z])"`)
)

func parseInput(input string) (rules map[int]Rule, messages []string) {
	parts := strings.Split(input, "\n\n")

	rules = make(map[int]Rule)
	ruleLines := strings.Split(parts[0], "\n")

	for _, line := range ruleLines {
		parts := strings.Split(line, ": ")
		id := util.MustAtoi(parts[0])
		if m := literalRe.FindStringSubmatch(parts[1]); len(m) > 0 {
			rules[id] = &Literal{char: m[1]}
		} else if strings.Contains(parts[1], "|") {
			leftRight := strings.Split(parts[1], " | ")
			rules[id] = &Choice{left: *NewSequenceFromString(leftRight[0]), right: *NewSequenceFromString(leftRight[1])}
		} else {
			rules[id] = NewSequenceFromString(parts[1])
		}
	}

	messages = strings.Split(parts[1], "\n")
	return rules, messages
}

func isSequenceMatch(rules *map[int]Rule, message string, rule Sequence) (bool, []int) {
	indexes := []int{0}
	for _, id := range rule.sequence {
		var newIndexes []int
		for _, idx := range indexes {
			match, lengths := isMatch(rules, message[idx:], id)
			if match {
				for _, l := range lengths {
					newIndexes = append(newIndexes, idx+l)
				}
			}
		}

		indexes = newIndexes
	}

	if len(indexes) > 0 {
		return true, indexes
	} else {
		return false, nil
	}
}

func isMatch(rules *map[int]Rule, message string, ruleId int) (matched bool, lengths []int) {
	if len(message) == 0 {
		return false, nil
	}

	rule := (*rules)[ruleId]
	switch rule.(type) {
	case *Literal:
		if rule.(*Literal).char == string(message[0]) {
			return true, []int{1}
		} else {
			return false, nil
		}
	case *Sequence:
		return isSequenceMatch(rules, message, *rule.(*Sequence))
	case *Choice:
		c := rule.(*Choice)
		matches := make([]int, 0)
		if match, indexes := isSequenceMatch(rules, message, c.left); match {
			matches = append(matches, indexes...)
		}
		if match, indexes := isSequenceMatch(rules, message, c.right); match {
			matches = append(matches, indexes...)
		}
		if len(matches) > 0 {
			return true, matches
		} else {
			return false, nil
		}
	}

	return false, nil
}

func Part1(rules map[int]Rule, messages []string) int {
	valid := 0
	for _, message := range messages {
		if match, lengths := isMatch(&rules, message, 0); match && util.IntInSlice(len(message), lengths) {
			valid++
		}
	}
	return valid
}

func Part2(rules map[int]Rule, messages []string) int {
	rules[8] = &Choice{
		left:  Sequence{sequence: []int{42}},
		right: Sequence{sequence: []int{42, 8}},
	}
	rules[11] = &Choice{
		left:  Sequence{sequence: []int{42, 31}},
		right: Sequence{sequence: []int{42, 11, 31}},
	}

	valid := 0
	for _, message := range messages {
		match, lengths := isMatch(&rules, message, 0)
		if match && util.IntInSlice(len(message), lengths) {
			valid++
		}
	}
	return valid
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day19.txt")
	rules, messages := parseInput(input)
	fmt.Println("(part1)", Part1(rules, messages)) //129
	fmt.Println("(part2)", Part2(rules, messages)) //243
}
