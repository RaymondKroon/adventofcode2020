package adventofcode2020

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func Atoi(s []string) ([]int, error) {
	var input []int
	for _, s := range s {
		i, _ := strconv.Atoi(s)
		input = append(input, i)
	}
	return input, nil
}

func MustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ReadInput(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(content), nil
}

func CreateSubstringSplitter(substr string) bufio.SplitFunc {
	splitter := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		// Return nothing if at end of file and no data passed
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := strings.Index(string(data), substr); i >= 0 {
			return i + len(substr), data[0:i], nil
		}

		// If at end of file with data return the data
		if atEOF {
			return len(data), data, nil
		}

		return
	}

	return splitter
}

func ReadSplittedInput(filePath string, splitOn string) ([]string, error) {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(CreateSubstringSplitter(splitOn))
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func ReadInputLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	defer file.Close()

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

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
