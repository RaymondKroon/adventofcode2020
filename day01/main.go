package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
)



func main() {
    file, err := os.Open("./input/day01.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var text []string

    for scanner.Scan() {
        text = append(text, scanner.Text())
    }

    var input []int
    for _, s := range text {
       i, _ := strconv.Atoi(s)
       input = append(input, i)
    }

    for iIdx, i := range input {
        for _, j := range input[iIdx+1:] {
            if i + j == 2020 {
                println(i*j)
            }
        }
    }



    for iIdx, i := range input {
        for jIdx, j := range input[iIdx+1:] {
            for _, k := range input[iIdx + jIdx + 1:] {
                if i+j+k == 2020 {
                    println(i * j * k)
                }
            }
        }
    }
}
