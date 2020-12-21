package main

import (
	"adventofcode2020/util"
	"fmt"
	"regexp"
	"strings"
)

type Rule interface {
	String() string
	IsMatch(message string) (matched bool, lengths []int)
}

type Literal struct {
	char string
}

type Sequence struct {
	sequence []int
	rules    *map[int]Rule
}

type Choice struct {
	left  Sequence
	right Sequence
}

func (l *Literal) String() string {
	return fmt.Sprintf("\"%s\"", l.char)
}

func (l *Literal) IsMatch(message string) (matched bool, lenght []int) {
	if len(message) == 0 {
		return false, nil
	}
	if l.char == string(message[0]) {
		return true, []int{1}
	} else {
		return false, nil
	}
}

func (s *Sequence) String() string {
	return strings.Join(util.MapIntsToStrings(s.sequence), " ")
}

func (s *Sequence) IsMatch(message string) (matched bool, lenght []int) {
	indexes := []int{0}
	for _, id := range s.sequence {
		var newIndexes []int
		for _, idx := range indexes {
			match, lengths := (*s.rules)[id].IsMatch(message[idx:])
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

func NewSequenceFromString(s string, rules *map[int]Rule) *Sequence {
	splitted := strings.Fields(s)
	seq := make([]int, len(splitted))

	for i, m := range splitted {
		seq[i] = util.MustAtoi(m)
	}

	return &Sequence{sequence: seq, rules: rules}
}

func (c *Choice) String() string {
	return fmt.Sprintf("%s | %s", c.left.String(), c.right.String())
}

func (c *Choice) IsMatch(message string) (matched bool, length []int) {
	matches := make([]int, 0)
	if match, indexes := c.left.IsMatch(message); match {
		matches = append(matches, indexes...)
	}
	if match, indexes := c.right.IsMatch(message); match {
		matches = append(matches, indexes...)
	}
	if len(matches) > 0 {
		return true, matches
	} else {
		return false, nil
	}
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
			rules[id] = &Choice{left: *NewSequenceFromString(leftRight[0], &rules), right: *NewSequenceFromString(leftRight[1], &rules)}
		} else {
			rules[id] = NewSequenceFromString(parts[1], &rules)
		}
	}

	messages = strings.Split(parts[1], "\n")
	return rules, messages
}

func Part1(rules map[int]Rule, messages []string) int {
	valid := 0
	for _, message := range messages {
		if match, lengths := rules[0].IsMatch(message); match && util.IntInSlice(len(message), lengths) {
			valid++
		}
	}
	return valid
}

func Part2(rules map[int]Rule, messages []string) int {
	rules[8] = &Choice{
		left:  Sequence{sequence: []int{42}, rules: &rules},
		right: Sequence{sequence: []int{42, 8}, rules: &rules},
	}
	rules[11] = &Choice{
		left:  Sequence{sequence: []int{42, 31}, rules: &rules},
		right: Sequence{sequence: []int{42, 11, 31}, rules: &rules},
	}

	valid := 0
	for _, message := range messages {
		match, lengths := rules[0].IsMatch(message)
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
