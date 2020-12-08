package main

import (
	"adventofcode2020"
	"container/list"
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

func selectBagWithColor(rules []BagRule, color string) BagRule {
	for _, br := range rules {
		if br.Color == color {
			return br
		}
	}
	return BagRule{}
}

func Part1CountShinyGoldParents(rules []BagRule) int {
	queue := list.New()
	queue.PushBack("shiny gold")
	selected := make([]string, 0)

	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)

		new := findBagsWichContains(rules, e.Value.(string))
		for _, n := range new {
			if !adventofcode2020.StringInSlice(n.Color, selected) {
				selected = append(selected, n.Color)
				queue.PushBack(n.Color)
			}

		}
	}

	return len(selected)
}

type q2 struct {
	color      string
	multiplier int
}

func Part2CountShinyGoldChildren(rules []BagRule) int {
	queue := list.New()
	queue.PushBack(q2{
		color:      "shiny gold",
		multiplier: 1,
	})

	total := 0

	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		val := e.Value.(q2)
		br := selectBagWithColor(rules, val.color)
		for _, b := range br.Bags {
			total += val.multiplier * b.N
			queue.PushBack(q2{
				color:      b.Color,
				multiplier: val.multiplier * b.N,
			})
		}
	}

	return total
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	inputLines, _ := adventofcode2020.ReadInputLines("./input/day07.txt")
	rules := ParseBagRules(inputLines)
	answer1 := Part1CountShinyGoldParents(rules) //172
	fmt.Printf("(part1) bags %d\n", answer1)

	answer2 := Part2CountShinyGoldChildren(rules) //39645
	fmt.Printf("(part2) bags %d\n", answer2)
}
