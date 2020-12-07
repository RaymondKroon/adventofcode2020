package main

import (
	"adventofcode2020"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ContainingBag struct {
	Color string
	N     int
}

type BagRule struct {
	Color string
	Bags  []ContainingBag
}

func ParseBagRules(input []string) []BagRule {
	colorRegex := regexp.MustCompile(`^([\w]+ [\w]+) bags contain`)
	containRegex := regexp.MustCompile(`(\d) ([\w]+ [\w]+) bag[s]?(?:, |.$)`)

	rules := make([]BagRule, 0)

	for _, l := range input {
		color := colorRegex.FindAllStringSubmatch(l, 1)[0][1]
		rule := BagRule{
			Color: color,
		}

		bags := make([]ContainingBag, 0)
		if strings.Contains(l, "no other bags") {

		} else {
			matches := containRegex.FindAllStringSubmatch(l, -1)
			for _, m := range matches {
				n, _ := strconv.Atoi(m[1])
				bags = append(bags, ContainingBag{
					Color: m[2],
					N:     n,
				})
			}
		}

		rule.Bags = bags

		rules = append(rules, rule)
	}

	return rules

}

func main() {
	inputLines, _ := adventofcode2020.ReadInputLines("./input/day07.txt")
	rules := ParseBagRules(inputLines)
	for _, r := range rules {
		fmt.Println(r)
	}
}
