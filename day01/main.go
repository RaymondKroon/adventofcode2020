package main

import (
    "adventofcode2020"
)

func main() {

    stringInput, _ := adventofcode2020.ReadInput("./input/day01.txt")
    input, _ := adventofcode2020.Atoi(stringInput)

    println(part1(input))
    println(part2(input))
}

func part1(input []int) int {
    for iIdx, i := range input {
        for _, j := range input[iIdx+1:] {
            if i + j == 2020 {
                return i*j
            }
        }
    }

    return -1
}

func part2(input []int) int {
    for iIdx, i := range input {
        for jIdx, j := range input[iIdx+1:] {
            if i+j < 2020 {
                for _, k := range input[iIdx+jIdx+1:] {
                    if i+j+k == 2020 {
                        return i * j * k
                    }
                }
            }
        }
    }
    return -1
}
