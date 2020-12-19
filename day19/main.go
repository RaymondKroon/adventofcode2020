package main

import (
	"adventofcode2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(input string) (rules map[int]string, messages []string) {
	parts := strings.Split(input, "\n\n")

	rules = make(map[int]string)
	ruleLines := strings.Split(parts[0], "\n")

	for _, line := range ruleLines {
		parts := strings.Split(line, ": ")
		id := util.MustAtoi(parts[0])
		rules[id] = parts[1]
	}

	messages = strings.Split(parts[1], "\n")
	return rules, messages
}

func expand(rules map[int]string, id int, depth int) string {
	rule := rules[id]
	expanded := ""
	if depth < 100 {
		for _, word := range strings.Fields(rule) {
			if i, err := strconv.Atoi(word); err == nil {
				expanded += expand(rules, i, depth+1)
			} else {
				expanded += word
			}
		}
	}

	return "(?:" + expanded + ")"
}

func Part1(rules map[int]string, messages []string) int {
	expandedRule := expand(rules, 0, 0)
	expandedRule = strings.NewReplacer("\"", "", " ", "").Replace(expandedRule)
	re := regexp.MustCompile(fmt.Sprint("^", expandedRule, "$"))
	valid := 0
	for _, message := range messages {
		if re.MatchString(message) {
			valid++
		}
	}
	return valid
}

func Part2(rules map[int]string, messages []string) int {
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"
	expandedRule := expand(rules, 0, 0)
	expandedRule = strings.NewReplacer("\"", "", " ", "").Replace(expandedRule)
	re := regexp.MustCompile(fmt.Sprint("^", expandedRule, "$"))
	valid := 0
	for _, message := range messages {
		if re.MatchString(message) {
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
	fmt.Println("(part1)", Part2(rules, messages)) //243
}
