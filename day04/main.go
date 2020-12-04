package main

import (
	"adventofcode2020"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParsePassports(input []string) *[]map[string]string {
	var passports []map[string]string
	p := make(map[string]string)

	for _, l := range input {
		if l == "" {
			passports = append(passports, p)
			p = make(map[string]string)
		} else {
			pairs := strings.Split(l, " ")
			for _, pair := range pairs {
				kv := strings.Split(pair, ":")
				p[string(kv[0])] = string(kv[1])
			}
		}
	}
	if len(p) > 0 {
		passports = append(passports, p)
	}

	return &passports
}

func PassportFieldsPresent(p map[string]string) bool {
	ok := true
	for _, k := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		_, exists := p[k]
		ok = ok && exists
	}

	return ok
}

func fieldBetweenIntValues(p map[string]string, field string, least int, most int) bool {
	fieldS := p[field]
	fieldI, err := strconv.Atoi(fieldS)
	if err != nil || fieldI < least || fieldI > most {
		return false
	}

	return true
}

func PassportFieldsPresentAndCorrect(p map[string]string) bool {
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

	hgt := p["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		hgtI, err := strconv.Atoi(hgt[:len(hgt)-2])
		if err != nil || hgtI < 150 || hgtI > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		hgtI, err := strconv.Atoi(hgt[:len(hgt)-2])
		if err != nil || hgtI < 59 || hgtI > 76 {
			return false
		}
	} else {
		return false
	}

	hcl := p["hcl"]
	matched, _ := regexp.MatchString("^#([a-fA-F0-9]{6})$", hcl)
	if !matched {
		return false
	}

	ecl := p["ecl"]
	if !adventofcode2020.StringInSlice(ecl, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}) {
		return false
	}

	pid := p["pid"]
	matched, _ = regexp.MatchString("^[0-9]{9}$", pid)
	if !matched {
		return false
	}

	return true
}

func main() {
	stringInput, _ := adventofcode2020.ReadInput("./input/day04.txt")
	passports := ParsePassports(stringInput)
	validPasswords := 0
	for _, p := range *passports {
		if PassportFieldsPresent(p) {
			validPasswords += 1
		}
	}

	fmt.Printf("(part1) Valid passwords: %d", validPasswords)

	validAndCorrectPasswords := 0
	for _, p := range *passports {
		if PassportFieldsPresentAndCorrect(p) {
			validAndCorrectPasswords += 1
		}
	}

	fmt.Printf("(part2) Valid passwords: %d", validAndCorrectPasswords)
}
