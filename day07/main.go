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

func findBagsWichContains(rules []BagRule, color string) []BagRule {
	result := make([]BagRule, 0)
	for _, r := range rules {
		for _, b := range r.Bags {
			if b.Color == color {
				result = append(result, r)
			}
		}
	}

	return result
}

func searchUp(rules []BagRule, color string, selected *[]string) {
	new := findBagsWichContains(rules, color)
	for _, br := range new {
		if !adventofcode2020.StringInSlice(br.Color, *selected) {
			*selected = append(*selected, br.Color)
			searchUp(rules, br.Color, selected)
		}
	}
	return
}

func selectBagWithColor(rules []BagRule, color string) BagRule {
	for _, br := range rules {
		if br.Color == color {
			return br
		}
	}
	return BagRule{}
}

func searchDown(rules []BagRule, color string) int {
	br := selectBagWithColor(rules, color)
	result := 0
	if len(br.Bags) > 0 {
		for _, b := range br.Bags {
			result += b.N * (1 + searchDown(rules, b.Color))
		}
	}
	return result
}

func Part1CountShinyGoldParents(rules []BagRule) int {
	total := make([]string, 0)
	searchUp(rules, "shiny gold", &total)

	return len(total)
}

func Part2CountShinyGoldChildren(rules []BagRule) int {
	return searchDown(rules, "shiny gold")
}

func main() {
	inputLines, _ := adventofcode2020.ReadInputLines("./input/day07.txt")
	rules := ParseBagRules(inputLines)
	answer1 := Part1CountShinyGoldParents(rules)
	fmt.Printf("(part1) bags %d\n", answer1)

	answer2 := Part2CountShinyGoldChildren(rules)
	fmt.Printf("(part2) bags %d\n", answer2)
}
