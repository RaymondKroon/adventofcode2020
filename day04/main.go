package main

import (
	"adventofcode2020"
	"fmt"
	"regexp"
	"strconv"
)

type Passport = map[string]string

func ParsePassports(chunks []string) *[]Passport {
	var passports []Passport
	passportRegex := regexp.MustCompile(`(\w{3}):(.+?)(\s|$)`)
	for _, c := range chunks {
		matches := passportRegex.FindAllStringSubmatch(c, -1)
		p := make(Passport)
		for _, kv := range matches {
			p[kv[1]] = kv[2]
		}

		passports = append(passports, p)
	}

	return &passports
}

func PassportFieldsPresent(p Passport) bool {
	ok := true
	for _, k := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		_, exists := p[k]
		ok = ok && exists
	}

	return ok
}

func fieldBetweenIntValues(p Passport, field string, least int, most int) bool {
	fieldS := p[field]
	fieldI, err := strconv.Atoi(fieldS)
	if err != nil || fieldI < least || fieldI > most {
		return false
	}

	return true
}

func PassportFieldsPresentAndCorrect(p Passport) bool {
	if !PassportFieldsPresent(p) {
		return false
	}

	if !fieldBetweenIntValues(p, "byr", 1920, 2002) {
		return false
	}

	if !fieldBetweenIntValues(p, "iyr", 2010, 2020) {
		return false
	}

	if !fieldBetweenIntValues(p, "eyr", 2020, 2030) {
		return false
	}

	hgtRegex := regexp.MustCompile(`^(\d+)(cm|in)$`)
	match := hgtRegex.FindStringSubmatch(p["hgt"])
	if len(match) == 0 {
		return false
	}
	switch match[2] {
	case "cm":
		if i, _ := strconv.Atoi(match[1]); i < 150 || i > 193 {
			return false
		}
	case "in":
		if i, _ := strconv.Atoi(match[1]); i < 59 || i > 76 {
			return false
		}
	default:
		return false
	}

	matched, _ := regexp.MatchString("^#([a-fA-F0-9]{6})$", p["hcl"])
	if !matched {
		return false
	}

	if !adventofcode2020.StringInSlice(p["ecl"], []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}) {
		return false
	}

	matched, _ = regexp.MatchString("^[0-9]{9}$", p["pid"])
	if !matched {
		return false
	}

	return true
}

func main() {
	defer adventofcode2020.Stopwatch("Run")()
	passwordChunks, _ := adventofcode2020.ReadSplittedInput("./input/day04.txt", "\n\n")
	passports := ParsePassports(passwordChunks)
	validPasswords := 0
	for _, p := range *passports {
		if PassportFieldsPresent(p) {
			validPasswords += 1
		}
	}

	fmt.Printf("(part1) Valid passwords: %d\n", validPasswords)

	validAndCorrectPasswords := 0
	for _, p := range *passports {
		if PassportFieldsPresentAndCorrect(p) {
			validAndCorrectPasswords += 1
		}
	}

	fmt.Printf("(part2) Valid passwords: %d\n", validAndCorrectPasswords)
}
