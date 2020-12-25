package main

import (
	"adventofcode2020/util"
	"fmt"
)

const subject = 7
const mod = 20201227

func findLoopSize(pk int) int {

	loop := 0
	for val := 1; val != pk; loop++ {
		val = val * 7 % mod
	}

	return loop
}

func calculateEncryptionKey(pkCard, pkDoor int) int {
	n := findLoopSize(pkCard)
	encryptionKey := 1
	for i := 0; i < n; i += 1 {
		encryptionKey = encryptionKey * pkDoor % mod
	}
	return encryptionKey
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInputLines("./input/day25.txt")
	card, door := util.MustAtoi(input[0]), util.MustAtoi(input[1])

	fmt.Println("(p1)", calculateEncryptionKey(card, door)) //11576351
}
