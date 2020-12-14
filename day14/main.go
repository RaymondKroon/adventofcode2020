package main

import (
	"adventofcode2020"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

type Mask struct {
	mask string
}

func (m Mask) apply(value int) int {
	for i, b := range m.mask {
		switch b {
		case 'X':
		case '0':
			value &= ^(1 << (35 - i))
		case '1':
			value |= 1 << (35 - i)
		}
	}

	return value
}

func (m Mask) applyV2(value int) int {
	for i, b := range m.mask {
		switch b {
		case 'X':
			panic("Permute first")
		case '0':
			value &= ^(1 << (35 - i))
		case '1':
			value |= 1 << (35 - i)
		}
	}

	return value
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func (m Mask) floatingMasks() []Mask {
	regex := regexp.MustCompile(`X`)
	matches := regex.FindAllStringIndex(m.mask, -1)
	indexes := make([]int, len(matches))
	for i, match := range matches {
		indexes[i] = match[0]
	}
	permutations := make([]Permutation, 0, 50)
	permute(indexes, []int{}, &permutations)
	masks := make([]Mask, len(permutations))
	for i, p := range permutations {
		mask := m.mask
		mask = strings.ReplaceAll(mask, "0", "-")
		for _, zero := range p.zeros {
			mask = replaceAtIndex(mask, '0', zero)
		}
		for _, one := range p.ones {
			mask = replaceAtIndex(mask, '1', one)
		}
		masks[i] = Mask{mask: mask}
	}
	return masks
}

type Permutation struct {
	zeros []int
	ones  []int
}

func contains(s []Permutation, e Permutation) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

func permute(zeros []int, ones []int, result *[]Permutation) {

	sort.Ints(zeros)
	sort.Ints(ones)

	if !contains(*result, Permutation{zeros: zeros, ones: ones}) {
		*result = append(*result, Permutation{zeros: zeros, ones: ones})
	}
	for i := 0; i < len(zeros); i++ {
		take := zeros[i]
		newZeros := make([]int, 0, len(zeros)-1)
		newZeros = append(newZeros, zeros[:i]...)
		newZeros = append(newZeros, zeros[i+1:]...)

		newOnes := make([]int, 0, len(ones)+1)
		newOnes = append(newOnes, ones...)
		newOnes = append(newOnes, take)
		permute(newZeros, newOnes, result)
	}
	return
}

type SetMem struct {
	address string
	value   int
}

func parseInstruction(input []string) []interface{} {
	result := make([]interface{}, len(input))
	instructionRegex := regexp.MustCompile(`([a-z\[\]0-9]+)\s=\s([0-9X]+)`)
	addressRegex := regexp.MustCompile(`[0-9]+`)

	for i, l := range input {
		m := instructionRegex.FindStringSubmatch(l)
		if m[1] == "mask" {
			result[i] = Mask{mask: m[2]}
		} else {
			address := addressRegex.FindStringSubmatch(m[1])[0]
			result[i] = SetMem{address: address, value: adventofcode2020.MustAtoi(m[2])}
		}
	}

	return result
}

func solvePart1(program []interface{}) int {
	mem := make(map[string]int)
	var mask Mask
	for _, instruction := range program {
		switch instruction.(type) {
		case Mask:
			mask = instruction.(Mask)
		case SetMem:
			sm := instruction.(SetMem)
			mem[sm.address] = mask.apply(sm.value)
		}
	}

	total := 0
	for _, v := range mem {
		if v != 0 {
			total += v
		}
	}

	return total
}

func solvePart2(program []interface{}) int {
	mem := make(map[int]int)
	var mask Mask
	var masks []Mask
	for _, instruction := range program {
		switch instruction.(type) {
		case Mask:
			mask = instruction.(Mask)
			masks = mask.floatingMasks()
		case SetMem:
			sm := instruction.(SetMem)
			//masks := mask.floatingMasks()
			for _, m := range masks {
				mem[m.applyV2(adventofcode2020.MustAtoi(sm.address))] = sm.value
			}
		}
	}

	total := 0
	for _, v := range mem {
		if v != 0 {
			total += v
		}
	}

	return total
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	lines, _ := adventofcode2020.ReadInputLines("./input/day14.txt")
	program := parseInstruction(lines)

	fmt.Println("(part1)", solvePart1(program)) // 10885823581193
	fmt.Println("(part2)", solvePart2(program)) // 3816594901962
}
