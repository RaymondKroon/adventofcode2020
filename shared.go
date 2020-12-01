package adventofcode2020

import (
    "bufio"
    "log"
    "os"
    "strconv"
)

func Atoi(s []string) ([]int, error) {
    var input []int
    for _, s := range s {
      i, _ := strconv.Atoi(s)
      input = append(input, i)
    }
    return input, nil
}

func ReadInput(filePath string) ([]string, error) {
    file, err := os.Open(filePath)

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var text []string

    for scanner.Scan() {
        text = append(text, scanner.Text())
    }

    return text, nil
}