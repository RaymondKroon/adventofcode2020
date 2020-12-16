package main

import (
	"adventofcode2020"
	"container/list"
	"fmt"
	"regexp"
	"strings"
)

const EOL = "\r\n"

type Range struct {
	Start        int
	IncludingEnd int
}

func (r *Range) InRange(value int) bool {
	return value >= r.Start && value <= r.IncludingEnd
}

type Rule struct {
	Name   string
	Ranges []Range
}

type FieldRange struct {
	inner map[string]bool
}

func CreateFieldRange() FieldRange {
	return FieldRange{inner: make(map[string]bool, 50)}
}

func (fr *FieldRange) First() string {
	for k, _ := range fr.inner {
		return k
	}

	panic("Empty FieldRange")
}

func (fr *FieldRange) Len() int {
	return len(fr.inner)
}

func (fr *FieldRange) Add(field string) {
	fr.inner[field] = true
}

func (fr *FieldRange) Remove(field string) {
	delete(fr.inner, field)
}

func (rule *Rule) ValueInRange(value int) bool {
	for _, _range := range rule.Ranges {
		if yes := _range.InRange(value); yes {
			return true
		}
	}

	return false
}

func (rule *Rule) ValuesInRange(values []int) bool {
	for _, v := range values {
		if ok := rule.ValueInRange(v); !ok {
			return false
		}
	}

	return true
}

func parseRules(input string) []Rule {
	regex := regexp.MustCompile(`([a-z\s]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)`)
	lines := strings.Split(input, "\n")
	rules := make([]Rule, len(lines))
	for i, line := range lines {
		match := regex.FindStringSubmatch(line)
		rules[i] = Rule{
			Name: match[1],
			Ranges: []Range{{
				Start:        adventofcode2020.MustAtoi(match[2]),
				IncludingEnd: adventofcode2020.MustAtoi(match[3]),
			},
				{
					Start:        adventofcode2020.MustAtoi(match[4]),
					IncludingEnd: adventofcode2020.MustAtoi(match[5]),
				}},
		}
	}

	return rules
}

func parseMyTicket(input string) []int {
	lines := strings.Split(input, EOL)
	nums := strings.Split(lines[1], ",")
	result, _ := adventofcode2020.Atoi(nums)
	return result
}

func parseOtherTickets(input string) [][]int {
	lines := strings.Split(input, EOL)
	result := make([][]int, len(lines)-1)
	for i, line := range lines[1:] {
		nums := strings.Split(line, ",")
		resultLine, _ := adventofcode2020.Atoi(nums)
		result[i] = resultLine
	}

	return result
}

func checkAllRule(rules []Rule) Rule {
	combinedRule := Rule{
		Name:   "checkAll",
		Ranges: make([]Range, 0),
	}
	for _, rule := range rules {
		combinedRule.Ranges = append(combinedRule.Ranges, rule.Ranges...)
	}

	return combinedRule
}

func Part1CheckErrorRate(tickets [][]int, rules []Rule) int {
	checkAll := checkAllRule(rules)

	invalid := []int{}
	for _, ticket := range tickets {
		for _, n := range ticket {
			if ok := checkAll.ValueInRange(n); !ok {
				invalid = append(invalid, n)
			}
		}
	}

	return adventofcode2020.Sum(invalid)
}

func Part2IdentifyValues(myTicket []int, otherTickets [][]int, rules []Rule) int {
	checkAll := checkAllRule(rules)

	nFields := len(myTicket)
	fieldOrder := make([]FieldRange, nFields)

	for i := 0; i < nFields; i++ {
		fieldOrder[i] = CreateFieldRange()
		for _, rule := range rules {
			fieldOrder[i].Add(rule.Name)
		}
	}

	for _, ticket := range otherTickets {
		if ok := checkAll.ValuesInRange(ticket); ok {
			for i, field := range ticket {
				for _, rule := range rules {
					if ok := rule.ValueInRange(field); !ok {
						fieldOrder[i].Remove(rule.Name)
					}
				}
			}
		}
	}

	singleFields := list.New()
	for _, fields := range fieldOrder {
		if fields.Len() == 1 {
			singleFields.PushBack(fields.First())
		}
	}

	for singleFields.Len() > 0 {
		e := singleFields.Front()
		singleFields.Remove(e)
		remove := e.Value.(string)
		for _, fo := range fieldOrder {
			if fo.Len() > 1 {
				fo.Remove(remove)
				if fo.Len() == 1 {
					singleFields.PushBack(fo.First())
				}
			}
		}
	}

	result := 1
	for i, fo := range fieldOrder {
		name := fo.First()

		if strings.HasPrefix(name, "departure") {
			result *= myTicket[i]
		}
	}

	return result
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	input, _ := adventofcode2020.ReadSplittedInput("./input/day16.txt", EOL+EOL)
	rules := parseRules(input[0])
	myTicket := parseMyTicket(input[1])
	otherTickets := parseOtherTickets(input[2])

	fmt.Println("(part1)", Part1CheckErrorRate(otherTickets, rules))
	fmt.Println("(part2)", Part2IdentifyValues(myTicket, otherTickets, rules))
}
