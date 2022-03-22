package main

import (
	"adventofcode2020/util"
	"fmt"
	"strings"
)

type IntArray []int

func (a IntArray) Equals(o IntArray) bool {
	if len(a) != len(o) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != o[i] {
			return false
		}
	}

	return true
}

func loadHands(input string) (player1 []int, player2 []int) {
	parts := strings.Split(input, "\n\n")
	for i, player := range []*[]int{&player1, &player2} {
		lines := strings.Split(parts[i], "\n")
		for _, c := range lines[1:] {
			*player = append(*player, util.MustAtoi(c))
		}
	}

	return
}

func playCombat(p1 []int, p2 []int) (winner int, score int) {
	hand1 := util.NewQueueFromSlice(p1)
	hand2 := util.NewQueueFromSlice(p2)

	for hand1.Len() > 0 && hand2.Len() > 0 {
		c1 := hand1.Pop()
		c2 := hand2.Pop()

		if *c1 > *c2 {
			hand1.PushBack(*c1)
			hand1.PushBack(*c2)
		} else {
			hand2.PushBack(*c2)
			hand2.PushBack(*c1)
		}
	}

	var winningHand util.Queue[int]
	if hand1.Len() > 0 {
		winner = 1
		winningHand = hand1
	} else {
		winner = 2
		winningHand = hand2
	}

	nCards := winningHand.Len()
	for i := nCards; i > 0; i-- {
		card := winningHand.Pop()
		score += i * *card
	}

	return
}

func playRecursiveCombat(p1, p2 []int) (winner int, score int) {
	hand1 := util.NewQueueFromSlice(p1)
	hand2 := util.NewQueueFromSlice(p2)

	var history1 []IntArray
	var history2 []IntArray
	recursiveBreak := false
	for hand1.Len() > 0 && hand2.Len() > 0 {

		if inSlice, _ := util.InSliceI[IntArray](hand1.Values(), history1); inSlice {
			recursiveBreak = true
			break
		}

		if inSlice, _ := util.InSliceI[IntArray](hand2.Values(), history2); inSlice {
			recursiveBreak = true
			break
		}

		history1 = append(history1, hand1.Values())
		history2 = append(history2, hand2.Values())
		c1 := hand1.Pop()
		c2 := hand2.Pop()

		roundWinner := -1

		if hand1.Len() >= *c1 && hand2.Len() >= *c2 {
			roundWinner, _ = playRecursiveCombat(hand1.Values()[:*c1], hand2.Values()[:*c2])
		}

		if roundWinner == 1 || (roundWinner < 0 && *c1 > *c2) {
			hand1.PushBack(*c1)
			hand1.PushBack(*c2)
		} else {
			hand2.PushBack(*c2)
			hand2.PushBack(*c1)
		}
	}

	var winningHand util.Queue[int]
	if recursiveBreak || hand1.Len() > 0 {
		winner = 1
		winningHand = hand1
	} else {
		winner = 2
		winningHand = hand2
	}

	nCards := winningHand.Len()
	for i := nCards; i > 0; i-- {
		card := winningHand.Pop()
		score += i * *card
	}

	return
}

func main() {
	defer util.Stopwatch("Run")()
	input, _ := util.ReadInput("./input/day22.txt")
	p1, p2 := loadHands(input)
	_, score1 := playCombat(p1, p2)
	fmt.Println("(part1)", score1) //32677

	_, score2 := playRecursiveCombat(p1, p2)
	fmt.Println("(part2)", score2) //33661

}
